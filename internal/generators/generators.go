package generators

import (
	"github.com/brianvoe/gofakeit/v7"
)

var Functions map[string]interface{} = map[string]interface{}{
    // File
    "CSV":            gofakeit.CSV,
    "JSON":           gofakeit.JSON,
    "XML":            gofakeit.XML,
    "FileExtension":  gofakeit.FileExtension,
    "FileMimeType":   gofakeit.FileMimeType,

    // Template
    "Template":   gofakeit.Template,
    "Markdown":   gofakeit.Markdown,
    "EmailText":  gofakeit.EmailText,
    "FixedWidth": gofakeit.FixedWidth,

    // ID
    "ID":   gofakeit.ID,
    "UUID": gofakeit.UUID,

    // Product
    "Product":            gofakeit.Product,
    "ProductName":        gofakeit.ProductName,
    "ProductDescription": gofakeit.ProductDescription,
    "ProductCategory":    gofakeit.ProductCategory,
    "ProductFeature":     gofakeit.ProductFeature,
    "ProductMaterial":    gofakeit.ProductMaterial,
    "ProductUPC":         gofakeit.ProductUPC,
    "ProductAudience":    gofakeit.ProductAudience,
    "ProductDimension":   gofakeit.ProductDimension,
    "ProductUseCase":     gofakeit.ProductUseCase,
    "ProductBenefit":     gofakeit.ProductBenefit,
    "ProductSuffix":      gofakeit.ProductSuffix,
    "ProductISBN":        gofakeit.ProductISBN,

    // Person
    "Person":         gofakeit.Person,
    "Name":           gofakeit.Name,
    "NamePrefix":     gofakeit.NamePrefix,
    "NameSuffix":     gofakeit.NameSuffix,
    "FirstName":      gofakeit.FirstName,
    "MiddleName":     gofakeit.MiddleName,
    "LastName":       gofakeit.LastName,
    "Gender":         gofakeit.Gender,
    "Age":            gofakeit.Age,
    "Ethnicity":      gofakeit.Ethnicity,
    "SSN":            gofakeit.SSN,
    "EIN":            gofakeit.EIN,
    "Hobby":          gofakeit.Hobby,
    "Contact":        gofakeit.Contact,
    "Email":          gofakeit.Email,
    "Phone":          gofakeit.Phone,
    "PhoneFormatted": gofakeit.PhoneFormatted,
    "Teams":          gofakeit.Teams,

    // Generate
    "Struct":   gofakeit.Struct,
    "Slice":    gofakeit.Slice,
    "Map":      gofakeit.Map,
    "Generate": gofakeit.Generate,
    "Regex":    gofakeit.Regex,

    // Auth
    "Username": gofakeit.Username,
    "Password": gofakeit.Password,

    // Address
    "Address":           gofakeit.Address,
    "City":              gofakeit.City,
    "Country":           gofakeit.Country,
    "CountryAbr":        gofakeit.CountryAbr,
    "State":             gofakeit.State,
    "StateAbr":          gofakeit.StateAbr,
    "Street":            gofakeit.Street,
    "StreetName":        gofakeit.StreetName,
    "StreetNumber":      gofakeit.StreetNumber,
    "StreetPrefix":      gofakeit.StreetPrefix,
    "StreetSuffix":      gofakeit.StreetSuffix,
    "Unit":              gofakeit.Unit,
    "Zip":               gofakeit.Zip,
    "Latitude":          gofakeit.Latitude,
    "LatitudeInRange":   gofakeit.LatitudeInRange,
    "Longitude":         gofakeit.Longitude,
    "LongitudeInRange":  gofakeit.LongitudeInRange,

    // Game
    "Gamertag": gofakeit.Gamertag,
    "Dice":     gofakeit.Dice,

    // Beer
    "BeerAlcohol": gofakeit.BeerAlcohol,
    "BeerBlg":     gofakeit.BeerBlg,
    "BeerHop":     gofakeit.BeerHop,
    "BeerIbu":     gofakeit.BeerIbu,
    "BeerMalt":    gofakeit.BeerMalt,
    "BeerName":    gofakeit.BeerName,
    "BeerStyle":   gofakeit.BeerStyle,
    "BeerYeast":   gofakeit.BeerYeast,

    // Car
    "Car":                   gofakeit.Car,
    "CarMaker":              gofakeit.CarMaker,
    "CarModel":              gofakeit.CarModel,
    "CarType":               gofakeit.CarType,
    "CarFuelType":           gofakeit.CarFuelType,
    "CarTransmissionType":   gofakeit.CarTransmissionType,

    // Words - Nouns
    "Noun":                   gofakeit.Noun,
    "NounCommon":             gofakeit.NounCommon,
    "NounConcrete":           gofakeit.NounConcrete,
    "NounAbstract":           gofakeit.NounAbstract,
    "NounCollectivePeople":   gofakeit.NounCollectivePeople,
    "NounCollectiveAnimal":   gofakeit.NounCollectiveAnimal,
    "NounCollectiveThing":    gofakeit.NounCollectiveThing,
    "NounCountable":          gofakeit.NounCountable,
    "NounUncountable":        gofakeit.NounUncountable,

    // Words - Verbs
    "Verb":         gofakeit.Verb,
    "VerbAction":   gofakeit.VerbAction,
    "VerbLinking":  gofakeit.VerbLinking,
    "VerbHelping":  gofakeit.VerbHelping,

    // Words - Adverbs
    "Adverb":                    gofakeit.Adverb,
    "AdverbManner":              gofakeit.AdverbManner,
    "AdverbDegree":              gofakeit.AdverbDegree,
    "AdverbPlace":               gofakeit.AdverbPlace,
    "AdverbTimeDefinite":        gofakeit.AdverbTimeDefinite,
    "AdverbTimeIndefinite":      gofakeit.AdverbTimeIndefinite,
    "AdverbFrequencyDefinite":   gofakeit.AdverbFrequencyDefinite,
    "AdverbFrequencyIndefinite": gofakeit.AdverbFrequencyIndefinite,

    // Words - Prepositions
    "Preposition":          gofakeit.Preposition,
    "PrepositionSimple":    gofakeit.PrepositionSimple,
    "PrepositionDouble":    gofakeit.PrepositionDouble,
    "PrepositionCompound":  gofakeit.PrepositionCompound,

    // Words - Adjectives
    "Adjective":                gofakeit.Adjective,
    "AdjectiveDescriptive":     gofakeit.AdjectiveDescriptive,
    "AdjectiveQuantitative":    gofakeit.AdjectiveQuantitative,
    "AdjectiveProper":          gofakeit.AdjectiveProper,
    "AdjectiveDemonstrative":   gofakeit.AdjectiveDemonstrative,
    "AdjectivePossessive":      gofakeit.AdjectivePossessive,
    "AdjectiveInterrogative":   gofakeit.AdjectiveInterrogative,
    "AdjectiveIndefinite":      gofakeit.AdjectiveIndefinite,

    // Words - Pronouns
    "Pronoun":               gofakeit.Pronoun,
    "PronounPersonal":       gofakeit.PronounPersonal,
    "PronounObject":         gofakeit.PronounObject,
    "PronounPossessive":     gofakeit.PronounPossessive,
    "PronounReflective":     gofakeit.PronounReflective,
    "PronounDemonstrative":  gofakeit.PronounDemonstrative,
    "PronounInterrogative":  gofakeit.PronounInterrogative,
    "PronounRelative":       gofakeit.PronounRelative,

    // Words - Connectives
    "Connective":            gofakeit.Connective,
    "ConnectiveTime":        gofakeit.ConnectiveTime,
    "ConnectiveComparative": gofakeit.ConnectiveComparative,
    "ConnectiveComplaint":   gofakeit.ConnectiveComplaint,
    "ConnectiveListing":     gofakeit.ConnectiveListing,
    "ConnectiveCasual":      gofakeit.ConnectiveCasual,
    "ConnectiveExamplify":   gofakeit.ConnectiveExamplify,

    // Words - General
    "Word":                   gofakeit.Word,
    "Sentence":               gofakeit.Sentence,
    "Paragraph":              gofakeit.Paragraph,
    "LoremIpsumWord":         gofakeit.LoremIpsumWord,
    "LoremIpsumSentence":     gofakeit.LoremIpsumSentence,
    "LoremIpsumParagraph":    gofakeit.LoremIpsumParagraph,
    "Question":               gofakeit.Question,
    "Quote":                  gofakeit.Quote,
    "Phrase":                 gofakeit.Phrase,

    // Foods
    "Fruit":     gofakeit.Fruit,
    "Vegetable": gofakeit.Vegetable,
    "Breakfast": gofakeit.Breakfast,
    "Lunch":     gofakeit.Lunch,
    "Dinner":    gofakeit.Dinner,
    "Snack":     gofakeit.Snack,
    "Dessert":   gofakeit.Dessert,

    // Misc
    "Bool":           gofakeit.Bool,
    "Weighted":       gofakeit.Weighted,
    "FlipACoin":      gofakeit.FlipACoin,
    "RandomMapKey":   gofakeit.RandomMapKey,
    "ShuffleAnySlice": gofakeit.ShuffleAnySlice,

    // Colors
    "Color":      gofakeit.Color,
    "HexColor":   gofakeit.HexColor,
    "RGBColor":   gofakeit.RGBColor,
    "SafeColor":  gofakeit.SafeColor,
    "NiceColors": gofakeit.NiceColors,

    // Images
    "Image":      gofakeit.Image,
    "ImageJpeg":  gofakeit.ImageJpeg,
    "ImagePng":   gofakeit.ImagePng,

    // Internet
    "URL":               gofakeit.URL,
    "UrlSlug":           gofakeit.UrlSlug,
    "DomainName":        gofakeit.DomainName,
    "DomainSuffix":      gofakeit.DomainSuffix,
    "IPv4Address":       gofakeit.IPv4Address,
    "IPv6Address":       gofakeit.IPv6Address,
    "MacAddress":        gofakeit.MacAddress,
    "HTTPStatusCode":    gofakeit.HTTPStatusCode,
    "HTTPStatusCodeSimple": gofakeit.HTTPStatusCodeSimple,
    "LogLevel":          gofakeit.LogLevel,
    "HTTPMethod":        gofakeit.HTTPMethod,
    "HTTPVersion":       gofakeit.HTTPVersion,
    "UserAgent":         gofakeit.UserAgent,
    "ChromeUserAgent":   gofakeit.ChromeUserAgent,
    "FirefoxUserAgent":  gofakeit.FirefoxUserAgent,
    "OperaUserAgent":    gofakeit.OperaUserAgent,
    "SafariUserAgent":   gofakeit.SafariUserAgent,
    "APIUserAgent":      gofakeit.APIUserAgent,

    // HTML
    "InputName": gofakeit.InputName,
    "Svg":       gofakeit.Svg,

    // Date/Time
    "Date":            gofakeit.Date,
    "PastDate":        gofakeit.PastDate,
    "FutureDate":      gofakeit.FutureDate,
    "DateRange":       gofakeit.DateRange,
    "NanoSecond":      gofakeit.NanoSecond,
    "Second":          gofakeit.Second,
    "Minute":          gofakeit.Minute,
    "Hour":            gofakeit.Hour,
    "Month":           gofakeit.Month,
    "MonthString":     gofakeit.MonthString,
    "Day":             gofakeit.Day,
    "WeekDay":         gofakeit.WeekDay,
    "Year":            gofakeit.Year,
    "TimeZone":        gofakeit.TimeZone,
    "TimeZoneAbv":     gofakeit.TimeZoneAbv,
    "TimeZoneFull":    gofakeit.TimeZoneFull,
    "TimeZoneOffset":  gofakeit.TimeZoneOffset,
    "TimeZoneRegion":  gofakeit.TimeZoneRegion,

    // Payment
    "Price":              gofakeit.Price,
    "CreditCard":         gofakeit.CreditCard,
    "CreditCardCvv":      gofakeit.CreditCardCvv,
    "CreditCardExp":      gofakeit.CreditCardExp,
    "CreditCardNumber":   gofakeit.CreditCardNumber,
    "CreditCardType":     gofakeit.CreditCardType,
    "Currency":           gofakeit.Currency,
    "CurrencyLong":       gofakeit.CurrencyLong,
    "CurrencyShort":      gofakeit.CurrencyShort,
    "AchRouting":         gofakeit.AchRouting,
    "AchAccount":         gofakeit.AchAccount,
    "BitcoinAddress":     gofakeit.BitcoinAddress,
    "BitcoinPrivateKey":  gofakeit.BitcoinPrivateKey,
    "BankName":           gofakeit.BankName,
    "BankType":           gofakeit.BankType,

    // Finance
    "Cusip": gofakeit.Cusip,
    "Isin":  gofakeit.Isin,

    // Company
    "BS":             gofakeit.BS,
    "Blurb":          gofakeit.Blurb,
    "BuzzWord":       gofakeit.BuzzWord,
    "Company":        gofakeit.Company,
    "CompanySuffix":  gofakeit.CompanySuffix,
    "Job":            gofakeit.Job,
    "JobDescriptor":  gofakeit.JobDescriptor,
    "JobLevel":       gofakeit.JobLevel,
    "JobTitle":       gofakeit.JobTitle,
    "Slogan":         gofakeit.Slogan,

    // Hacker
    "HackerAbbreviation": gofakeit.HackerAbbreviation,
    "HackerAdjective":    gofakeit.HackerAdjective,
    "Hackeringverb":      gofakeit.HackeringVerb,
    "HackerNoun":         gofakeit.HackerNoun,
    "HackerPhrase":       gofakeit.HackerPhrase,
    "HackerVerb":         gofakeit.HackerVerb,

    // Hipster
    "HipsterWord":      gofakeit.HipsterWord,
    "HipsterSentence":  gofakeit.HipsterSentence,
    "HipsterParagraph": gofakeit.HipsterParagraph,

    // App
    "AppName":    gofakeit.AppName,
    "AppVersion": gofakeit.AppVersion,
    "AppAuthor":  gofakeit.AppAuthor,

    // Animal
    "PetName":    gofakeit.PetName,
    "Animal":     gofakeit.Animal,
    "AnimalType": gofakeit.AnimalType,
    "FarmAnimal": gofakeit.FarmAnimal,
    "Cat":        gofakeit.Cat,
    "Dog":        gofakeit.Dog,
    "Bird":       gofakeit.Bird,

    // Emoji
    "Emoji":            gofakeit.Emoji,
    "EmojiCategory":    gofakeit.EmojiCategory,
    "EmojiAlias":       gofakeit.EmojiAlias,
    "EmojiTag":         gofakeit.EmojiTag,
    "EmojiFlag":        gofakeit.EmojiFlag,
    "EmojiAnimal":      gofakeit.EmojiAnimal,
    "EmojiFood":        gofakeit.EmojiFood,
    "EmojiPlant":       gofakeit.EmojiPlant,
    "EmojiMusic":       gofakeit.EmojiMusic,
    "EmojiVehicle":     gofakeit.EmojiVehicle,
    "EmojiSport":       gofakeit.EmojiSport,
    "EmojiFace":        gofakeit.EmojiFace,
    "EmojiHand":        gofakeit.EmojiHand,
    "EmojiClothing":    gofakeit.EmojiClothing,
    "EmojiLandmark":    gofakeit.EmojiLandmark,
    "EmojiElectronics": gofakeit.EmojiElectronics,
    "EmojiGame":        gofakeit.EmojiGame,
    "EmojiTools":       gofakeit.EmojiTools,
    "EmojiWeather":     gofakeit.EmojiWeather,
    "EmojiJob":         gofakeit.EmojiJob,
    "EmojiPerson":      gofakeit.EmojiPerson,
    "EmojiGesture":     gofakeit.EmojiGesture,
    "EmojiCostume":     gofakeit.EmojiCostume,
    "EmojiSentence":    gofakeit.EmojiSentence,

    // Language
    "Language":             gofakeit.Language,
    "LanguageAbbreviation": gofakeit.LanguageAbbreviation,
    "ProgrammingLanguage":  gofakeit.ProgrammingLanguage,

    // Number
    "Number":         gofakeit.Number,
    "Int":            gofakeit.Int,
    "IntN":           gofakeit.IntN,
    "Int8":           gofakeit.Int8,
    "Int16":          gofakeit.Int16,
    "Int32":          gofakeit.Int32,
    "Int64":          gofakeit.Int64,
    "Uint":           gofakeit.Uint,
    "UintN":          gofakeit.UintN,
    "Uint8":          gofakeit.Uint8,
    "Uint16":         gofakeit.Uint16,
    "Uint32":         gofakeit.Uint32,
    "Uint64":         gofakeit.Uint64,
    "Float32":        gofakeit.Float32,
    "Float32Range":   gofakeit.Float32Range,
    "Float64":        gofakeit.Float64,
    "Float64Range":   gofakeit.Float64Range,
    "ShuffleInts":    gofakeit.ShuffleInts,
    "RandomInt":      gofakeit.RandomInt,
    "HexUint":        gofakeit.HexUint,

    // String
    "Digit":          gofakeit.Digit,
    "DigitN":         gofakeit.DigitN,
    "Letter":         gofakeit.Letter,
    "LetterN":        gofakeit.LetterN,
    "Lexify":         gofakeit.Lexify,
    "Numerify":       gofakeit.Numerify,
    "ShuffleStrings": gofakeit.ShuffleStrings,
    "RandomString":   gofakeit.RandomString,

    // Celebrity
    "CelebrityActor":   gofakeit.CelebrityActor,
    "CelebrityBusiness": gofakeit.CelebrityBusiness,
    "CelebritySport":   gofakeit.CelebritySport,

    // Minecraft
    "MinecraftOre":            gofakeit.MinecraftOre,
    "MinecraftWood":           gofakeit.MinecraftWood,
    "MinecraftArmorTier":      gofakeit.MinecraftArmorTier,
    "MinecraftArmorPart":      gofakeit.MinecraftArmorPart,
    "MinecraftWeapon":         gofakeit.MinecraftWeapon,
    "MinecraftTool":           gofakeit.MinecraftTool,
    "MinecraftDye":            gofakeit.MinecraftDye,
    "MinecraftFood":           gofakeit.MinecraftFood,
    "MinecraftAnimal":         gofakeit.MinecraftAnimal,
    "MinecraftVillagerJob":    gofakeit.MinecraftVillagerJob,
    "MinecraftVillagerStation": gofakeit.MinecraftVillagerStation,
    "MinecraftVillagerLevel":  gofakeit.MinecraftVillagerLevel,
    "MinecraftMobPassive":     gofakeit.MinecraftMobPassive,
    "MinecraftMobNeutral":     gofakeit.MinecraftMobNeutral,
    "MinecraftMobHostile":     gofakeit.MinecraftMobHostile,
    "MinecraftMobBoss":        gofakeit.MinecraftMobBoss,
    "MinecraftBiome":          gofakeit.MinecraftBiome,
    "MinecraftWeather":        gofakeit.MinecraftWeather,

    // Book
    "Book":       gofakeit.Book,
    "BookTitle":  gofakeit.BookTitle,
    "BookAuthor": gofakeit.BookAuthor,
    "BookGenre":  gofakeit.BookGenre,

    // Movie
    "Movie":      gofakeit.Movie,
    "MovieName":  gofakeit.MovieName,
    "MovieGenre": gofakeit.MovieGenre,

    // Error
    "Error":             gofakeit.Error,
    "ErrorDatabase":     gofakeit.ErrorDatabase,
    "ErrorGRPC":         gofakeit.ErrorGRPC,
    "ErrorHTTP":         gofakeit.ErrorHTTP,
    "ErrorHTTPClient":   gofakeit.ErrorHTTPClient,
    "ErrorHTTPServer":   gofakeit.ErrorHTTPServer,
    "ErrorRuntime":      gofakeit.ErrorRuntime,

    // School
    "School": gofakeit.School,

    // Song
    "Song":       gofakeit.Song,
    "SongName":   gofakeit.SongName,
    "SongArtist": gofakeit.SongArtist,
    "SongGenre":  gofakeit.SongGenre,
}
