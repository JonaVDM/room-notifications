package rooms

type Thing struct {
	Data     []Datum  `json:"data"`
	Metadata Metadata `json:"_metadata"`
}

type Datum struct {
	ID                                    int64                                   `json:"ID"`
	DatumID                               int64                                   `json:"id"`
	ExtraInformatieURL                    interface{}                             `json:"extraInformatieUrl"`
	HuurtoeslagMogelijk                   int64                                   `json:"huurtoeslagMogelijk"`
	Postalcode                            string                                  `json:"postalcode"`
	Street                                string                                  `json:"street"`
	HouseNumber                           string                                  `json:"houseNumber"`
	HouseNumberAddition                   string                                  `json:"houseNumberAddition"`
	GemeenteGeoLocatieNaam                GemeenteGeoLocatieNaam                  `json:"gemeenteGeoLocatieNaam"`
	RentBuy                               string                                  `json:"rentBuy"`
	Vhe                                   string                                  `json:"VHE"`
	EnergyIndex                           interface{}                             `json:"energyIndex"`
	Zonnepanelen                          interface{}                             `json:"zonnepanelen"`
	GaslozeWoning                         interface{}                             `json:"gaslozeWoning"`
	NulOpDeMeterWoning                    interface{}                             `json:"nulOpDeMeterWoning"`
	AvailableFromDate                     string                                  `json:"availableFromDate"`
	TotalRent                             float64                                 `json:"totalRent"`
	NetRent                               float64                                 `json:"netRent"`
	CalculationRent                       float64                                 `json:"calculationRent"`
	ServiceCosts                          float64                                 `json:"serviceCosts"`
	HeatingCosts                          float64                                 `json:"heatingCosts"`
	Epv                                   interface{}                             `json:"epv"`
	AdditionalCosts                       float64                                 `json:"additionalCosts"`
	EenmaligeKosten                       int64                                   `json:"eenmaligeKosten"`
	FlexibelHurenActief                   int64                                   `json:"flexibelHurenActief"`
	SellingPrice                          interface{}                             `json:"sellingPrice"`
	Latitude                              float64                                 `json:"latitude"`
	Longitude                             float64                                 `json:"longitude"`
	ReactieURL                            string                                  `json:"reactieUrl"`
	NewlyBuild                            int64                                   `json:"newlyBuild"`
	AreaDwelling                          int64                                   `json:"areaDwelling"`
	VolumeDwelling                        interface{}                             `json:"volumeDwelling"`
	AreaPerceel                           interface{}                             `json:"areaPerceel"`
	AreaLivingRoom                        int64                                   `json:"areaLivingRoom"`
	AreaSleepingRoom                      string                                  `json:"areaSleepingRoom"`
	Lengte                                interface{}                             `json:"lengte"`
	Breedte                               interface{}                             `json:"breedte"`
	Hoogte                                interface{}                             `json:"hoogte"`
	Tuin                                  int64                                   `json:"tuin"`
	StorageRoom                           int64                                   `json:"storageRoom"`
	ActionLabelFrom                       interface{}                             `json:"actionLabelFrom"`
	ActionLabelUntil                      interface{}                             `json:"actionLabelUntil"`
	Balcony                               int64                                   `json:"balcony"`
	ConstructionYear                      string                                  `json:"constructionYear"`
	VatInclusive                          int64                                   `json:"vatInclusive"`
	AantalMedebewoners                    *int64                                  `json:"aantalMedebewoners"`
	ToewijzingModelTypeInCode             string                                  `json:"toewijzingModelTypeInCode"`
	IsZelfstandig                         int64                                   `json:"isZelfstandig"`
	VerzameladvertentieID                 interface{}                             `json:"verzameladvertentieID"`
	MinimumIncome                         interface{}                             `json:"minimumIncome"`
	MaximumIncome                         interface{}                             `json:"maximumIncome"`
	MinimumHouseholdSize                  interface{}                             `json:"minimumHouseholdSize"`
	MaximumHouseholdSize                  interface{}                             `json:"maximumHouseholdSize"`
	MinimumAge                            int64                                   `json:"minimumAge"`
	MaximumAge                            interface{}                             `json:"maximumAge"`
	InwonendeKinderenMinimum              interface{}                             `json:"inwonendeKinderenMinimum"`
	InwonendeKinderenMaximum              interface{}                             `json:"inwonendeKinderenMaximum"`
	LeeftijdControleVoorBeideAanvrager    int64                                   `json:"leeftijdControleVoorBeideAanvrager"`
	GebruikAangepasteHuurinkomenstabel    int64                                   `json:"gebruikAangepasteHuurinkomenstabel"`
	VoorrangLeeftijdMin                   interface{}                             `json:"voorrangLeeftijdMin"`
	VoorrangLeeftijdMax                   interface{}                             `json:"voorrangLeeftijdMax"`
	VoorrangHuishoudgrootteMin            interface{}                             `json:"voorrangHuishoudgrootteMin"`
	VoorrangHuishoudgrootteMax            interface{}                             `json:"voorrangHuishoudgrootteMax"`
	VoorrangGezinnenKinderen              *int64                                  `json:"voorrangGezinnenKinderen"`
	VoorrangKernbinding                   interface{}                             `json:"voorrangKernbinding"`
	VoorrangAlleUrgenties                 int64                                   `json:"voorrangAlleUrgenties"`
	ExtraInschrijfduurUitgeschakeld       int64                                   `json:"extraInschrijfduurUitgeschakeld"`
	ToewijzingID                          int64                                   `json:"toewijzingID"`
	NumberOfReactions                     int64                                   `json:"numberOfReactions"`
	ToewijzingHeeftAnnuleerReden          int64                                   `json:"toewijzingHeeftAnnuleerReden"`
	IsExtraAanbod                         int64                                   `json:"isExtraAanbod"`
	ReagerenZonderInschrijvingMogelijk    string                                  `json:"reagerenZonderInschrijvingMogelijk"`
	HuurinkomenstabelGebruiken            int64                                   `json:"huurinkomenstabelGebruiken"`
	ModelHuurinkomenstabelActief          int64                                   `json:"modelHuurinkomenstabelActief"`
	ModelInZoekprofiel                    int64                                   `json:"modelInZoekprofiel"`
	ShowEnergyCosts                       int64                                   `json:"showEnergyCosts"`
	WoningTypeInZoekprofiel               int64                                   `json:"woningTypeInZoekprofiel"`
	PublicationDate                       string                                  `json:"publicationDate"`
	ClosingDate                           string                                  `json:"closingDate"`
	ToewijzingIsGeannuleerd               int64                                   `json:"toewijzingIsGeannuleerd"`
	IsWoningruil                          int64                                   `json:"isWoningruil"`
	Regio                                 Regio                                   `json:"regio"`
	Land                                  City                                    `json:"land"`
	Municipality                          City                                    `json:"municipality"`
	City                                  City                                    `json:"city"`
	Neighborhood                          interface{}                             `json:"neighborhood"`
	Quarter                               Quarter                                 `json:"quarter"`
	Corporation                           Corporation                             `json:"corporation"`
	DwellingType                          DwellingType                            `json:"dwellingType"`
	EnergyLabel                           *HuurtoeslagVoorwaarde                  `json:"energyLabel"`
	Heating                               Heating                                 `json:"heating"`
	SleepingRoom                          SleepingRoom                            `json:"sleepingRoom"`
	Kitchen                               Heating                                 `json:"kitchen"`
	Attic                                 interface{}                             `json:"attic"`
	OppervlakteTuin                       interface{}                             `json:"oppervlakteTuin"`
	Floor                                 Floor                                   `json:"floor"`
	RentDuration                          interface{}                             `json:"rentDuration"`
	ActionLabel                           interface{}                             `json:"actionLabel"`
	Sorteergroep                          Sorteergroep                            `json:"sorteergroep"`
	HuurtoeslagVoorwaarde                 HuurtoeslagVoorwaarde                   `json:"huurtoeslagVoorwaarde"`
	GardenSite                            interface{}                             `json:"gardenSite"`
	BalconySite                           interface{}                             `json:"balconySite"`
	Woningsoort                           HuurtoeslagVoorwaarde                   `json:"woningsoort"`
	ToewijzingModelCategorie              HuurtoeslagVoorwaarde                   `json:"toewijzingModelCategorie"`
	Model                                 Model                                   `json:"model"`
	Koopvoorwaarden                       interface{}                             `json:"koopvoorwaarden"`
	SpecifiekeVoorzieningen               []SpecifiekeVoorzieningen               `json:"specifiekeVoorzieningen"`
	Doelgroepen                           []HuurtoeslagVoorwaarde                 `json:"doelgroepen"`
	ServicecomponentenBinnenServicekosten []ServicecomponentenBinnenServicekosten `json:"servicecomponentenBinnenServicekosten"`
	ServicecomponentenBuitenServicekosten []interface{}                           `json:"servicecomponentenBuitenServicekosten"`
	Bestemming                            interface{}                             `json:"bestemming"`
	Toegankelijkheid                      []interface{}                           `json:"toegankelijkheid"`
	EenmaligeKostenSpecificatie           []interface{}                           `json:"eenmaligeKostenSpecificatie"`
	Floorplans                            []Floorplan                             `json:"floorplans"`
	Pictures                              []Picture                               `json:"pictures"`
	EnergyCosts                           []interface{}                           `json:"energyCosts"`
	Videos                                []interface{}                           `json:"videos"`
	URLKey                                string                                  `json:"urlKey"`
	InschrijvingVereistVoorReageren       bool                                    `json:"inschrijvingVereistVoorReageren"`
	IsExternModelType                     bool                                    `json:"isExternModelType"`
	ObjectType                            string                                  `json:"objectType"`
	SorteringBeleidsregels                []interface{}                           `json:"sorteringBeleidsregels"`
	ReactieBeleidsregels                  []ReactieBeleidsregel                   `json:"reactieBeleidsregels"`
	RelatieHuurInkomenRegels              []RelatieHuurInkomenRegel               `json:"relatieHuurInkomenRegels"`
	ReactionData                          ReactionData                            `json:"reactionData"`
	Woningruilprofiel                     Woningruilprofiel                       `json:"woningruilprofiel"`
	ActionLabelIfActive                   bool                                    `json:"actionLabelIfActive"`
	ActionLabelIsActive                   bool                                    `json:"actionLabelIsActive"`
	AvailableFrom                         string                                  `json:"availableFrom"`
	Infoveld                              string                                  `json:"infoveld"`
	InfoveldKort                          string                                  `json:"infoveldKort"`
	InfoveldBewoners                      string                                  `json:"infoveldBewoners"`
	ActielabelToelichting                 string                                  `json:"actielabelToelichting"`
	IsGepubliceerd                        bool                                    `json:"isGepubliceerd"`
}

type City struct {
	ID   int64                  `json:"id"`
	Name GemeenteGeoLocatieNaam `json:"name"`
}

type Corporation struct {
	ID               int64  `json:"id"`
	Code             int64  `json:"code"`
	Logo             string `json:"logo"`
	HuurtoeslagTonen int64  `json:"huurtoeslagTonen"`
	Name             string `json:"name"`
	LocalizedName    string `json:"localizedName"`
}

type HuurtoeslagVoorwaarde struct {
	ID                int64   `json:"id"`
	Code              *Code   `json:"code,omitempty"`
	Icon              string  `json:"icon"`
	LocalizedNaam     *string `json:"localizedNaam,omitempty"`
	LocalizedIconText *string `json:"localizedIconText,omitempty"`
	ToonOpWebsite     *int64  `json:"toonOpWebsite,omitempty"`
	IsZelfstandig     *int64  `json:"isZelfstandig,omitempty"`
}

type DwellingType struct {
	ID                  int64     `json:"id"`
	HuurprijsDuurActief int64     `json:"huurprijsDuurActief"`
	Categorie           Categorie `json:"categorie"`
	Code                string    `json:"code"`
	Name                string    `json:"name"`
	LocalizedName       string    `json:"localizedName"`
}

type Floor struct {
	ID            int64  `json:"id"`
	Verdieping    int64  `json:"verdieping"`
	LocalizedName string `json:"localizedName"`
}

type Floorplan struct {
	Label string `json:"label"`
	URI   string `json:"uri"`
}

type Heating struct {
	ID            int64  `json:"id"`
	LocalizedName string `json:"localizedName"`
}

type Model struct {
	ID                                int64                 `json:"id"`
	Code                              string                `json:"code"`
	AdvertentieSluitenNaEersteReactie int64                 `json:"advertentieSluitenNaEersteReactie"`
	EinddatumTonen                    int64                 `json:"einddatumTonen"`
	AantalReactiesTonen               int64                 `json:"aantalReactiesTonen"`
	SlaagkansTonen                    int64                 `json:"slaagkansTonen"`
	MotiverenBijReagerenActief        int64                 `json:"motiverenBijReagerenActief"`
	ModelCategorie                    HuurtoeslagVoorwaarde `json:"modelCategorie"`
	IsHospiteren                      bool                  `json:"isHospiteren"`
}

type Picture struct {
	Label string `json:"label"`
	Type  Type   `json:"type"`
	URI   string `json:"uri"`
}

type Quarter struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	PlaatsID int64  `json:"plaatsId"`
}

type ReactieBeleidsregel struct {
	ID               int64             `json:"id"`
	Code             string            `json:"code"`
	Parameters       []interface{}     `json:"parameters"`
	ParametersAsText Woningruilprofiel `json:"parametersAsText"`
}

type Woningruilprofiel struct {
}

type ReactionData struct {
	Loggedin                  bool    `json:"loggedin"`
	AangepasteNettoHuurprijs  float64 `json:"aangepasteNettoHuurprijs"`
	AangepasteTotaleHuurprijs float64 `json:"aangepasteTotaleHuurprijs"`
	RechtOpHuurprijsverlaging bool    `json:"rechtOpHuurprijsverlaging"`
	ZoekprofielMatch          bool    `json:"zoekprofielMatch"`
	IsPassend                 bool    `json:"isPassend"`
	KanReageren               bool    `json:"kanReageren"`
	ZoekprofielMatchOrder     int64   `json:"zoekprofielMatchOrder"`
	ToonAlsVoorrangsgroep     bool    `json:"toonAlsVoorrangsgroep"`
}

type Regio struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	PostcodesMetVoorrang string `json:"postcodesMetVoorrang"`
}

type RelatieHuurInkomenRegel struct {
	CorporatieID              int64           `json:"corporatieId"`
	PlaatsID                  interface{}     `json:"plaatsId"`
	HuishoudGrootte           HuishoudGrootte `json:"huishoudGrootte"`
	Huurprijs                 HuishoudGrootte `json:"huurprijs"`
	Inkomen                   HuishoudGrootte `json:"inkomen"`
	AowLeeftijdAlsLeeftijdMin int64           `json:"aowLeeftijdAlsLeeftijdMin"`
	AowLeeftijdAlsLeeftijdMax int64           `json:"aowLeeftijdAlsLeeftijdMax"`
	Leeftijd                  HuishoudGrootte `json:"leeftijd"`
	InkomenformVerplicht      int64           `json:"inkomenformVerplicht"`
	InkomenUitzondering       HuishoudGrootte `json:"inkomenUitzondering"`
	InkomensGroepID           int64           `json:"inkomensGroepId"`
}

type HuishoudGrootte struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type ServicecomponentenBinnenServicekosten struct {
	ID            int64         `json:"id"`
	LocalizedNaam LocalizedNaam `json:"localizedNaam"`
}

type SleepingRoom struct {
	ID            int64  `json:"id"`
	AmountOfRooms int64  `json:"amountOfRooms"`
	LocalizedName string `json:"localizedName"`
	Naam          string `json:"naam"`
}

type Sorteergroep struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type SpecifiekeVoorzieningen struct {
	ID                   int64     `json:"id"`
	InCode               string    `json:"inCode"`
	DwellingTypeCategory Categorie `json:"dwellingTypeCategory"`
	LocalizedLabel       string    `json:"localizedLabel"`
	LocalizedName        string    `json:"localizedName"`
}

type Metadata struct {
	Page             int64 `json:"page"`
	Limit            int64 `json:"limit"`
	PageCount        int64 `json:"page_count"`
	TotalSearchCount int64 `json:"total_search_count"`
	TotalCount       int64 `json:"total_count"`
}

type GemeenteGeoLocatieNaam string

const (
	Groningen GemeenteGeoLocatieNaam = "Groningen"
	Nederland GemeenteGeoLocatieNaam = "Nederland"
)

type Code string

const (
	Hospitereninschrijfduur Code = "hospitereninschrijfduur"
	Inschrijfduur           Code = "inschrijfduur"
	Jongeren                Code = "jongeren"
)

type Categorie string

const (
	Woning Categorie = "woning"
)

type Type string

const (
	Jpg Type = "jpg"
)

type LocalizedNaam string

const (
	BelastingenEnHeffingen LocalizedNaam = "Belastingen en heffingen"
	Elektriciteit          LocalizedNaam = "Elektriciteit"
	Glasverzekering        LocalizedNaam = "Glasverzekering"
	Verwarming             LocalizedNaam = "Verwarming"
	Water                  LocalizedNaam = "Water"
)
