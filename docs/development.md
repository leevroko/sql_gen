# Development

This is a document that is dedicated to logging the development process and plans.

# Entries 

## 26-04-28 

I want to implement several features that are going to be mentioned in [[Development plans]].

The one that I am going to focus on right now is Related generation.

> [!info] ## Related generation development plan 
> 
> You can define relations between columns in different tables that allow the generator to create relations between entries. Related columns must have the same type
> 
> ```yaml
> # ...
> tables: 
> - name: tasks
>   require_entry_count: 100
>   fields: 
>     - name: task_id
>     # ...
>     - name: employee_id
>       type: int
>       generated_type: Uint16
>   relations:
>     - our_column: employee_id
>       their_table: employees
>       their_column: emplyee_id
>       relation_type: A_includes_B # B_includes_A, A_is_B, A_intersects_B
>       difference_coefficient: 0.3
> # ...
> ```
> 
> - `our_column` - a column from the table where the relation is defined
> - `their_table` - a table which is related to `our_table`
> - `their_column` - a column from `their_table`
> - `relation_type` - a type of relation between the data in the tables, where A is `our_table`, B is `their_table`:
>     - `A_includes_B` - means that A references all the data from B, but not all rows from A reference B
>     - `B_includes_A` - means that every row from A references some row from B, but not all B rows are referenced
>     - `A_is_B` - means that A references all the data from B and every row from A references some row from B
>     - `A_intersects_B` - means that some A rows reference some B rows, but not all A rows refernce B and not all B rows are referenced
> - `difference_coefficient` - from 0 to 1, means slightly different things in every relation type, meaningless in `A_is_B` (equal to 0). Is equal to number of rows that do not have references divided by all rows in the table, which has rows, that don't have references.
> - `disparity_coefficient` - from 0 to 1, influences the distribution of references. Higher disparity means some rows get more references, while others have less. Lower disparity means more equal distribution of references.

Well, this is how I see this feature, but it will not be implemented like this right now.

For now I will only implement relation types and difference coefficient.

First of all, we need to have in mind, that generating relations creates a depndency on availability of the data that is required to be referenced. This creates a problem, when the table is being generated earlier than the one it references. In this case, we have to have some kind of data cache to use for pre generation of referenced data on demand, and its reusal on actual table creation.

The structure will be as follows:
- A head map that stores key-table name and value-table map pairs 
- Table maps that store key-column name and value-data slice of slices pairs 
- A slice of slices for chunked data generation

This structure is faulty because if we generate a big reference cache, we can run out of RAM eventually. To address this issue I suggest creating service-tables in the database during data generation. The more general approach is to use hard drive to store caches, but as it is a complex task in on its own, reusing the DB for that seems like a plausible solution.

Now the structure should be like this:
- A head map that stores key-table name and value-table map pairs 
- Table maps that store key-column name and value-service table name pairs 
- Mono column service tables that store pregenerated data

So now we have to consider the cache existance when generating a table and we should also consider the `difference_coefficient` when considering the cache usage. But no, we can just count rows in a service table and subtract the count from the required row count and then we will have the number of rows, that are still required to be generated, if any.

But wait, we don't need to generate separate tables for storing the pregenerated data, because it can be generated in chunks as we thought before and stored in the same table that references it. Then we just need to place the reference of the original table's column instead of the specially generated table's column.

To summarize, what we have now.

### Stages of table generation

1. Reading the configuration and table schema
2. Checking cache for pre generated data 
    1. Considering the difference between the cached data count and required count
3. Considering the required relations 
    1. Creating an entry in the cache map for columnt in our table that has the pregenerated reference data 
    2. Generating the referenced data table rows as we go 
4. Generating the actual table data and using the referenced data table(-s) if needed
5. Deleting the cache entry and pregenerated data from the DB for the current table, if it was referenced by any other table before

### Cache structure 

```go 
type ServiceTableName string
type ColumnName string
type TableName string 
type TableCacheInfo map[ColumnName]ServiceTableName

type Cache struct {
    storage map[TableName]TableCacheInfo
}
```

### Plans update 

Now before we proceed any further, we need to adress the chunked generation feature as implementing it after the related generation will be hard and carry risks of breaking things. We need to define the required interfaces and control flow right now so that the foundation is in place when we need it. 

Chunked generation should consider the hardware it is running on, data it is going to generate and calculate a maximum query size for the task. It should use this query size to do multiple steps of query generation, query execution to ensure correct RAM usage and complete data generation.

```go 
if entriesAsked > d.entryLimit {
	panic(fmt.Sprintf("...", d.entryLimit, entriesAsked))
}

stmts := d.createQueries(tableSchema, entriesAsked)

for _, stmt := range stmts {
	cmdTags, err := d.db.Exec(ctx, stmt)
	if err != nil {
        // ...
		return err
	}
    // ...
}
return nil

```

As we can see in the code, the query generation is done all at once, even though as a slice of queries. 

But wait, there is a serious problem in our task definition. We can not generate queries concurrently as generating multiple queries at once can result in RAM overflow by our definition. So we do not need to consider concurrency and just need to do the generation iteratively.

So the updated code should look something like this:

TODO
