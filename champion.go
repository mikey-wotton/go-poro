package go_poro

import (
	"log"
	"regexp"
	"strings"

	league "github.com/mikey-wotton/go-league"
)

type Champion struct {
	Name ChampName
	Tags []*Tag
}

type ChampName league.ChampionName

func (c ChampName) Valid() (bool, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	champName, err := c.ToURI()
	if err != nil {
		return false, err
	}

	champName = strings.ToLower(champName)

	for _, championName := range validChampNames {
		name := strings.ToLower(reg.ReplaceAllString(string(championName), ""))
		if name == champName {
			return true, nil
		}
	}

	return false, nil
}

func (c ChampName) ToURI() (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(string(c), ""), nil
}


var validChampNames = []league.ChampionName{
	league.ChampionNameUnknown,
	league.ChampionNameAatrox,
	league.ChampionNameAhri,
	league.ChampionNameAkali,
	league.ChampionNameAlistar,
	league.ChampionNameAmumu,
	league.ChampionNameAnivia,
	league.ChampionNameAnnie,
	league.ChampionNameAphelios,
	league.ChampionNameAshe,
	league.ChampionNameAurelion,
	league.ChampionNameAzir,
	league.ChampionNameBard,
	league.ChampionNameBlitzcrank,
	league.ChampionNameBrand,
	league.ChampionNameBraum,
	league.ChampionNameCaitlyn,
	league.ChampionNameCamille,
	league.ChampionNameCassiopeia,
	league.ChampionNameChoGath,
	league.ChampionNameCorki,
	league.ChampionNameDarius,
	league.ChampionNameDiana,
	league.ChampionNameDrMundo,
	league.ChampionNameDraven,
	league.ChampionNameEkko,
	league.ChampionNameElise,
	league.ChampionNameEvelynn,
	league.ChampionNameEzreal,
	league.ChampionNameFiddlesticks,
	league.ChampionNameFiora,
	league.ChampionNameFizz,
	league.ChampionNameGalio,
	league.ChampionNameGangplank,
	league.ChampionNameGaren,
	league.ChampionNameGnar,
	league.ChampionNameGragas,
	league.ChampionNameGraves,
	league.ChampionNameHecarim,
	league.ChampionNameHeimerdinger,
	league.ChampionNameIllaoi,
	league.ChampionNameIrelia,
	league.ChampionNameIvern,
	league.ChampionNameJanna,
	league.ChampionNameJarvan,
	league.ChampionNameJax,
	league.ChampionNameJayce,
	league.ChampionNameJhin,
	league.ChampionNameJinx,
	league.ChampionNameKaiSa,
	league.ChampionNameKalista,
	league.ChampionNameKarma,
	league.ChampionNameKarthus,
	league.ChampionNameKassadin,
	league.ChampionNameKatarina,
	league.ChampionNameKayle,
	league.ChampionNameKayn,
	league.ChampionNameKennen,
	league.ChampionNameKhaZix,
	league.ChampionNameKindred,
	league.ChampionNameKled,
	league.ChampionNameKogMaw,
	league.ChampionNameLeBlanc,
	league.ChampionNameLeeSin,
	league.ChampionNameLeona,
	league.ChampionNameLillia,
	league.ChampionNameLissandra,
	league.ChampionNameLucian,
	league.ChampionNameLulu,
	league.ChampionNameLux,
	league.ChampionNameMalphite,
	league.ChampionNameMalzahar,
	league.ChampionNameMaokai,
	league.ChampionNameMasterYi,
	league.ChampionNameMissFortune,
	league.ChampionNameMordekaiser,
	league.ChampionNameMorgana,
	league.ChampionNameNami,
	league.ChampionNameNasus,
	league.ChampionNameNautilus,
	league.ChampionNameNeeko,
	league.ChampionNameNidalee,
	league.ChampionNameNocturne,
	league.ChampionNameNunuAndWillump,
	league.ChampionNameOlaf,
	league.ChampionNameOrianna,
	league.ChampionNameOrnn,
	league.ChampionNamePantheon,
	league.ChampionNamePoppy,
	league.ChampionNamePyke,
	league.ChampionNameQiyana,
	league.ChampionNameQuinn,
	league.ChampionNameRakan,
	league.ChampionNameRammus,
	league.ChampionNameRekSai,
	league.ChampionNameRell,
	league.ChampionNameRenekton,
	league.ChampionNameRengar,
	league.ChampionNameRiven,
	league.ChampionNameRumble,
	league.ChampionNameRyze,
	league.ChampionNameSamira,
	league.ChampionNameSejuani,
	league.ChampionNameSenna,
	league.ChampionNameSeraphine,
	league.ChampionNameSett,
	league.ChampionNameShaco,
	league.ChampionNameShen,
	league.ChampionNameShyvana,
	league.ChampionNameSinged,
	league.ChampionNameSion,
	league.ChampionNameSivir,
	league.ChampionNameSkarner,
	league.ChampionNameSona,
	league.ChampionNameSoraka,
	league.ChampionNameSwain,
	league.ChampionNameSylas,
	league.ChampionNameSyndra,
	league.ChampionNameTahmKench,
	league.ChampionNameTaliyah,
	league.ChampionNameTalon,
	league.ChampionNameTaric,
	league.ChampionNameTeemo,
	league.ChampionNameThresh,
	league.ChampionNameTristana,
	league.ChampionNameTrundle,
	league.ChampionNameTryndamere,
	league.ChampionNameTwistedFate,
	league.ChampionNameTwitch,
	league.ChampionNameUdyr,
	league.ChampionNameUrgot,
	league.ChampionNameVarus,
	league.ChampionNameVayne,
	league.ChampionNameVeigar,
	league.ChampionNameVelKoz,
	league.ChampionNameVi,
	league.ChampionNameViktor,
	league.ChampionNameVladimir,
	league.ChampionNameVolibear,
	league.ChampionNameWarwick,
	league.ChampionNameWukong,
	league.ChampionNameXayah,
	league.ChampionNameXerath,
	league.ChampionNameXinZhao,
	league.ChampionNameYasuo,
	league.ChampionNameYone,
	league.ChampionNameYorick,
	league.ChampionNameYuumi,
	league.ChampionNameZac,
	league.ChampionNameZed,
	league.ChampionNameZiggs,
	league.ChampionNameZilean,
	league.ChampionNameZoe,
	league.ChampionNameZyra,
}