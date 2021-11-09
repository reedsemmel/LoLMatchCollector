#!/usr/bin/env python3

from playground import item_id_name, mythics
from dataclasses import dataclass

champion_id_to_name = {
    "266": "Aatrox",
    "103": "Ahri",
    "84": "Akali",
    "166": "Akshan",
    "12": "Alistar",
    "32": "Amumu",
    "34": "Anivia",
    "1": "Annie",
    "523": "Aphelios",
    "22": "Ashe",
    "136": "AurelionSol",
    "268": "Azir",
    "432": "Bard",
    "53": "Blitzcrank",
    "63": "Brand",
    "201": "Braum",
    "51": "Caitlyn",
    "164": "Camille",
    "69": "Cassiopeia",
    "31": "Chogath",
    "42": "Corki",
    "122": "Darius",
    "131": "Diana",
    "119": "Draven",
    "36": "DrMundo",
    "245": "Ekko",
    "60": "Elise",
    "28": "Evelynn",
    "81": "Ezreal",
    "9": "FiddleSticks",
    "114": "Fiora",
    "105": "Fizz",
    "3": "Galio",
    "41": "Gangplank",
    "86": "Garen",
    "150": "Gnar",
    "79": "Gragas",
    "104": "Graves",
    "887": "Gwen",
    "120": "Hecarim",
    "74": "Heimerdinger",
    "420": "Illaoi",
    "39": "Irelia",
    "427": "Ivern",
    "40": "Janna",
    "59": "JarvanIV",
    "24": "Jax",
    "126": "Jayce",
    "202": "Jhin",
    "222": "Jinx",
    "145": "Kaisa",
    "429": "Kalista",
    "43": "Karma",
    "30": "Karthus",
    "38": "Kassadin",
    "55": "Katarina",
    "10": "Kayle",
    "141": "Kayn",
    "85": "Kennen",
    "121": "Khazix",
    "203": "Kindred",
    "240": "Kled",
    "96": "KogMaw",
    "7": "Leblanc",
    "64": "LeeSin",
    "89": "Leona",
    "876": "Lillia",
    "127": "Lissandra",
    "236": "Lucian",
    "117": "Lulu",
    "99": "Lux",
    "54": "Malphite",
    "90": "Malzahar",
    "57": "Maokai",
    "11": "MasterYi",
    "21": "MissFortune",
    "62": "MonkeyKing",
    "82": "Mordekaiser",
    "25": "Morgana",
    "267": "Nami",
    "75": "Nasus",
    "111": "Nautilus",
    "518": "Neeko",
    "76": "Nidalee",
    "56": "Nocturne",
    "20": "Nunu",
    "2": "Olaf",
    "61": "Orianna",
    "516": "Ornn",
    "80": "Pantheon",
    "78": "Poppy",
    "555": "Pyke",
    "246": "Qiyana",
    "133": "Quinn",
    "497": "Rakan",
    "33": "Rammus",
    "421": "RekSai",
    "526": "Rell",
    "58": "Renekton",
    "107": "Rengar",
    "92": "Riven",
    "68": "Rumble",
    "13": "Ryze",
    "360": "Samira",
    "113": "Sejuani",
    "235": "Senna",
    "147": "Seraphine",
    "875": "Sett",
    "35": "Shaco",
    "98": "Shen",
    "102": "Shyvana",
    "27": "Singed",
    "14": "Sion",
    "15": "Sivir",
    "72": "Skarner",
    "37": "Sona",
    "16": "Soraka",
    "50": "Swain",
    "517": "Sylas",
    "134": "Syndra",
    "223": "TahmKench",
    "163": "Taliyah",
    "91": "Talon",
    "44": "Taric",
    "17": "Teemo",
    "412": "Thresh",
    "18": "Tristana",
    "48": "Trundle",
    "23": "Tryndamere",
    "4": "TwistedFate",
    "29": "Twitch",
    "77": "Udyr",
    "6": "Urgot",
    "110": "Varus",
    "67": "Vayne",
    "45": "Veigar",
    "161": "Velkoz",
    "711": "Vex",
    "254": "Vi",
    "234": "Viego",
    "112": "Viktor",
    "8": "Vladimir",
    "106": "Volibear",
    "19": "Warwick",
    "498": "Xayah",
    "101": "Xerath",
    "5": "XinZhao",
    "157": "Yasuo",
    "777": "Yone",
    "83": "Yorick",
    "350": "Yuumi",
    "154": "Zac",
    "238": "Zed",
    "115": "Ziggs",
    "26": "Zilean",
    "142": "Zoe",
    "143": "Zyra",
}

base_mythics = [
    "",
    "2065",
    "3068",
    "3078",
    "3152",
    "3190",
    "4005",
    "4633",
    "4636",
    "6617",
    "6630",
    "6631",
    "6632",
    "6653",
    "6655",
    "6656",
    "6662",
    "6664",
    "6671",
    "6672",
    "6673",
    "6691",
    "6692",
    "6693",
]


@dataclass
class TableEntry:
    champion: str
    mythic: str
    totalGames: int
    wins: int
    losses: int
    totalKills: int
    totalDeaths: int
    totalAssists: int
    gamesTop: int
    gamesJungle: int
    gamesMid: int
    gamesCarry: int
    gamesUtility: int
    gamesUnknown: int

def initialize() -> 'dict[dict[TableEntry]]':
    table: 'dict[dict[TableEntry]]' = dict()
    for _, champion in champion_id_to_name.items():
        table[champion] = {}
        for mythic_id in base_mythics:
            table[champion][item_id_name[mythic_id]] = TableEntry(
                champion=champion,
                mythic=item_id_name[mythic_id],
                totalGames=0,
                wins=0,
                losses=0,
                totalKills=0,
                totalDeaths=0,
                totalAssists=0,
                gamesTop=0,
                gamesJungle=0,
                gamesMid=0,
                gamesCarry=0,
                gamesUtility=0,
                gamesUnknown=0,
            )
        table[champion]["None"] = TableEntry(
            champion=champion,
            mythic="None",
            totalGames=0,
            wins=0,
            losses=0,
            totalKills=0,
            totalDeaths=0,
            totalAssists=0,
            gamesTop=0,
            gamesJungle=0,
            gamesMid=0,
            gamesCarry=0,
            gamesUtility=0,
            gamesUnknown=0,
        )
    return table

@dataclass
class MatchEntry:
    championName: str
    mythicItem: str
    win: str
    kills: str
    deaths: str
    assists: str
    teamPosition: str

def get_all_matches():
    import sqlite3

    def nt_factory(cursor, row):
        return MatchEntry(*row)

    table = initialize()
    conn = sqlite3.connect('matches.sqlite3')
    conn.row_factory = nt_factory
    cur = conn.cursor()
    i = 0
    for row in cur.execute("SELECT championName, mythicItem, win, kills, deaths, assists, teamPosition FROM matches;"):
        if i % 100000 == 0:
            print("Currently on", i)
        i += 1
        update_table(table, row)
    conn.close()

    return table



def update_table(table: 'dict[dict[TableEntry]]', match: 'MatchEntry') -> None:
    ent = table[match.championName][item_id_name[match.mythicItem]]
    ent.totalGames += 1
    if match.win == "true":
        ent.wins += 1
    else:
        ent.losses += 1
    ent.totalKills += int(match.kills)
    ent.totalDeaths += int(match.deaths)
    ent.totalAssists += int(match.assists)
    if match.teamPosition == "TOP":
        ent.gamesTop += 1
    elif match.teamPosition == "JUNGLE":
        ent.gamesJungle += 1
    elif match.teamPosition == "MIDDLE":
        ent.gamesMid += 1
    elif match.teamPosition == "BOTTOM":
        ent.gamesCarry += 1
    elif match.teamPosition == "UTILITY":
        ent.gamesUtility += 1
    else:
        ent.gamesUnknown += 1

def add_aggregates(table: 'dict[dict[TableEntry]]') -> None:

    champ_aggregates: 'dict[TableEntry]' = dict()
    mythic_aggregates: 'dict[TableEntry]' = dict()
    for mythic_id in base_mythics:
        mythic_aggregates[item_id_name[mythic_id]] = TableEntry(
            champion="All",
            mythic=item_id_name[mythic_id],
            totalGames=0,
            wins=0,
            losses=0,
            totalKills=0,
            totalDeaths=0,
            totalAssists=0,
            gamesTop=0,
            gamesJungle=0,
            gamesMid=0,
            gamesCarry=0,
            gamesUtility=0,
            gamesUnknown=0,
        )

    for _, champion in champion_id_to_name.items():
        champ_aggregates[champion] = TableEntry(
            champion=champion,
            mythic="All",
            totalGames=0,
            wins=0,
            losses=0,
            totalKills=0,
            totalDeaths=0,
            totalAssists=0,
            gamesTop=0,
            gamesJungle=0,
            gamesMid=0,
            gamesCarry=0,
            gamesUtility=0,
            gamesUnknown=0,
        )

    for champion, entries in table.items():
        for mythic, ent in entries.items():
            champ_aggregates[champion].totalGames += ent.totalGames
            champ_aggregates[champion].wins += ent.wins
            champ_aggregates[champion].losses += ent.losses
            champ_aggregates[champion].totalKills += ent.totalKills
            champ_aggregates[champion].totalDeaths += ent.totalDeaths
            champ_aggregates[champion].totalAssists += ent.totalAssists
            champ_aggregates[champion].gamesTop += ent.gamesTop
            champ_aggregates[champion].gamesJungle += ent.gamesJungle
            champ_aggregates[champion].gamesMid += ent.gamesMid
            champ_aggregates[champion].gamesCarry += ent.gamesCarry
            champ_aggregates[champion].gamesUtility += ent.gamesUtility
            champ_aggregates[champion].gamesUnknown += ent.gamesUnknown

            mythic_aggregates[mythic].totalGames += ent.totalGames
            mythic_aggregates[mythic].wins += ent.wins
            mythic_aggregates[mythic].losses += ent.losses
            mythic_aggregates[mythic].totalKills += ent.totalKills
            mythic_aggregates[mythic].totalDeaths += ent.totalDeaths
            mythic_aggregates[mythic].totalAssists += ent.totalAssists
            mythic_aggregates[mythic].gamesTop += ent.gamesTop
            mythic_aggregates[mythic].gamesJungle += ent.gamesJungle
            mythic_aggregates[mythic].gamesMid += ent.gamesMid
            mythic_aggregates[mythic].gamesCarry += ent.gamesCarry
            mythic_aggregates[mythic].gamesUtility += ent.gamesUtility
            mythic_aggregates[mythic].gamesUnknown += ent.gamesUnknown

    return champ_aggregates, mythic_aggregates

def insert_data(table, champ_ag, mythic_ag):
    import sqlite3

    query = """INSERT INTO aggregate(totalGames,wins,losses,totalKills,totalDeaths,totalAssists,
    gamesTop,gamesJungle,gamesMid,gamesCarry,gamesUtility,gamesUnknown,champion,mythic)
    VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);"""

    conn = sqlite3.connect('aggregates.sqlite3')
    cur = conn.cursor()
    for _, entries in table.items():
        for _, ent in entries.items():
            cur.execute(query, (
                ent.totalGames,
                ent.wins,
                ent.losses,
                ent.totalKills,
                ent.totalDeaths,
                ent.totalAssists,
                ent.gamesTop,
                ent.gamesJungle,
                ent.gamesMid,
                ent.gamesCarry,
                ent.gamesUtility,
                ent.gamesUnknown,
                ent.champion,
                ent.mythic,
            ))

    for _, ent in champ_ag.items():
        cur.execute(query, (
            ent.totalGames,
            ent.wins,
            ent.losses,
            ent.totalKills,
            ent.totalDeaths,
            ent.totalAssists,
            ent.gamesTop,
            ent.gamesJungle,
            ent.gamesMid,
            ent.gamesCarry,
            ent.gamesUtility,
            ent.gamesUnknown,
            ent.champion,
            ent.mythic,
        ))

    for _, ent in mythic_ag.items():
        cur.execute(query, (
            ent.totalGames,
            ent.wins,
            ent.losses,
            ent.totalKills,
            ent.totalDeaths,
            ent.totalAssists,
            ent.gamesTop,
            ent.gamesJungle,
            ent.gamesMid,
            ent.gamesCarry,
            ent.gamesUtility,
            ent.gamesUnknown,
            ent.champion,
            ent.mythic,
        ))

    conn.commit()
    conn.close()



if __name__ == "__main__":
    table = get_all_matches()
    champ_aggregates, mythic_aggregates = add_aggregates(table)
    insert_data(table, champ_aggregates, mythic_aggregates)
