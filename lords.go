package main

type Lord struct {
	Number      int
	Race        string
	Faction     string
	LordName    string
	Description string
}

const (
	// races
	Beastmen      = "Beastmen"
	Brettonia     = "Bretonnia"
	DemonsOfChaos = "Demons of Chaos"
	DarkElves     = "Dark Elves"
	Dwarfs        = "Dwarfs"
	TheEmpire     = "The Empire"
	GrandCathay   = "Grand Cathay"
	Greenskins    = "Greenskins"
	HighElves     = "HighElves"

	placeholder = "placeholder"
)

func createLords() []Lord {
	var l []Lord

	l = append(l, newLord(Beastmen, "Warherd of the One-Eye", "Khazrak the One-Eye", "Khazrak terrorizes the Empire's northern forests, making Blood-grounds of the Drakwald once more."))
	l = append(l, newLord(Beastmen, "Harbinger of Disaster", "Malagor the Dark Omen", "Malagor revels in the ruins of the Badlands to remind his followers of what will come to pass if they fulfil his desires for destruction."))
	l = append(l, newLord(Beastmen, "Warherd of the Shadowgave", "Morghur the Shadowgave", "Reborn once more, Morghur's Ruinous cycle of madness and destruction begins anew in the forests of Estalia."))
	l = append(l, newLord(Beastmen, "Slaugterhorn Tribe", "Taurox the Brass Bull", "Taurox's latest rampage has led him on a stampede through Naggaroth to Clar Karond, where begin his new Blood-Grounds."))
	l = append(l, newLord(Brettonia, "Couronne", "Louen Leoncouer", "King Louen rides out from Couronne with the blessings of the Lady to eradicate evil from Bretonnia's lands and beyond."))
	l = append(l, newLord(Brettonia, "Bordeleaux Errant", "Alberic de Bordeleaux", "The Duke of Bordeleaux embarks upon his long-desired Grail Quest from the Bretonnian colony of Bregonne on the coast of Lustria."))
	l = append(l, newLord(Brettonia, "Carcassonne", "The Fay Enchantress", "The Fay Enchantress channels the Lady to wield powerful magic against all threats to her home of Carcassonne."))
	l = append(l, newLord(Brettonia, "Chevaliers de Lyonesse", "Repanse de Lyonesse", "Repanse walks in the footsteps of ancient Errantry Knights amoungst the Coast of Araby's haunted dunes, smiting Bretonnia's enemies in the Lady's name."))
	l = append(l, newLord(DemonsOfChaos, "Legion of Chaos", "Daniel", "The Demon Prince enters the Forest of Decay in the northern Chaos Wastes, where he will compete with Archaon himself for the ultimate leadership of Chaos."))
	l = append(l, newLord(DarkElves, "Naggarond", "Malekith", "In his capital Naggarond, Malekith has brought his rebellious Dreadlords to heel and now sets about bringing the entire world under his dominion."))
	l = append(l, newLord(DarkElves, "Cult of Pleasure", "Morathi", "Banished to the Ancient City of Quintex, Morathis works her plans for world domination far south from the prying eyes of Naggarond."))
	l = append(l, newLord(DarkElves, "Har Ganeth", "Crone Hellebrone", "Har Ganeth, the city of Khaine, seeks war, and thus Hellebron goes forth to wage it and demand the world's submission!"))
	l = append(l, newLord(DarkElves, "The Blessed Dread", "Lokhir Fellheart", "The Krakenlord has sailed to Haichai on the coastline of distant Cathay, seeking to plunder the wealth and riches of this exotic land."))
	l = append(l, newLord(DarkElves, "Hag Graef", "Malus Darkblade", "In search of the true Warpsword of Khaine, Malus Darkblade wanders through the maddening Chaos Wastes to the Black Rock."))
	l = append(l, newLord(DarkElves, "The Thousand Maws", "Rakarth", "The Beastlord has reached the Great Turtle Isle off the coast of Lustria in search of ever greater monsters to enslave to his will."))
	l = append(l, newLord(Dwarfs, "Karaz-a-Karak", "Thorgrim Grudgebearer", "King Thorgrim continues his Age of Reckoning, setting out to crush those with unanswered transgressions starting with those in the shadow of the Everpeak itself."))
	l = append(l, newLord(Dwarfs, "Karak Kadrin", "Ungrim Ironfist", "Ungrim Ironfist returns to Slayer Keep to put to death all who threaten the sanctity of Kark Kadrin and the surrounding Dwarfholds."))
	l = append(l, newLord(Dwarfs, "Clan Angrund", "Belegar Ironhammer", "Belegar plots his campaign to reclaim his birthright, the kingdom of Karak Eight Peaks, from the mountain Hold of Zarakzil."))
	l = append(l, newLord(Dwarfs, "The Ancestral Throng", "Grombrindal - The White Dwarf", "The White Dwarf ventures to the Drakla Spire in Naggaroth, seeking to settle a grudge with the Lord of Naggarond almost as ancient as his own beard."))
	l = append(l, newLord(Dwarfs, "Ironbrow's Expedition", "Thorek Ironbrow", "After reclaiming the lost Dwarfhold of Karak Zorn, Thorek continues his great task to recover the lost relics of his kind."))
	l = append(l, newLord(TheEmpire, "Reikland", "Emperor Karl Franz", placeholder))
	l = append(l, newLord(TheEmpire, "The Golden Order", "Balthasar Gelt", placeholder))
	l = append(l, newLord(TheEmpire, "The Huntmarshal's Expedition", "Markus Wulfhart", placeholder))
	l = append(l, newLord(TheEmpire, "Cult of Sigmar", "Volkmar the Grim", placeholder))
	l = append(l, newLord(GrandCathay, "The Northern Provinces", "Miao Ying, the Storm Dragon", placeholder))
	l = append(l, newLord(GrandCathay, "The Western Provinces", "Zhao Ming, the Iron Dragon", placeholder))
	l = append(l, newLord(Greenskins, "Grimgor's 'Ardboyz", "Grimgor Ironhide", placeholder))
	l = append(l, newLord(Greenskins, "Bonerattlaz", "Azhag the Slaughterer", placeholder))
	l = append(l, newLord(Greenskins, "Crooked Moon", "Skarsnik", placeholder))
	l = append(l, newLord(Greenskins, "The Bloody Handz", "Wurrzag Da Great Green Prophet", placeholder))
	l = append(l, newLord(Greenskins, "Broken Axe", "Grom the Paunch", placeholder))
	l = append(l, newLord(HighElves, "Eataine", "Tyrion", placeholder))
	l = append(l, newLord(HighElves, "Order of Loremasters", "Teclis", placeholder))
	l = append(l, newLord(HighElves, "Avelorn", "Alarielle the Radiant", placeholder))
	l = append(l, newLord(HighElves, "Nagarythe", "Alith Anar", placeholder))
	

	// adding an index afterwards is quicker than typing
	for i, lord := range l {
		lord.Number = i
	}

	return l
}

func newLord(race string, faction string, name string, description string) Lord {
	var l Lord
	l.Race = race
	l.Faction = faction
	l.LordName = name
	l.Description = description
	return l

}
