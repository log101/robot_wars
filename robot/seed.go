package robot

// SEED DATA
// SKILLS
var Strike = Skill{
	name:       "strike",
	hpEffect:   -50,
	critChance: 10,
}

var Slam = Skill{
	name:       "slam",
	hpEffect:   -20,
	critChance: 20,
}

var Purge = Skill{
	name:       "purge",
	hpEffect:   -30,
	critChance: 30,
}

var LesserHeal = Skill{
	name:       "lesser heal",
	hpEffect:   20,
	critChance: 0,
}

var GreaterHeal = Skill{
	name:       "greater heal",
	hpEffect:   50,
	critChance: 0,
}

var Splint = Skill{
	name:       "splint",
	hpEffect:   -45,
	critChance: 40,
}

var Destruction = Skill{
	name:       "destruction",
	hpEffect:   -70,
	critChance: 15,
}

var Immobilize = Skill{
	name:       "immobilize",
	hpEffect:   -10,
	critChance: 20,
}

var Slice = Skill{
	name:       "slice",
	hpEffect:   -5,
	critChance: 90,
}

var Rampage = Skill{
	name:       "rampage",
	hpEffect:   -90,
	critChance: 0,
}

// ROBOTS
var Tera = Robot{
	id:     1,
	name:   "Tera",
	health: 100,
	skill1: Rampage,
	skill2: Slice,
	skill3: Immobilize,
}

var Imotron = Robot{
	id:     1,
	name:   "Imotron",
	health: 100,
	skill1: Strike,
	skill2: Slam,
	skill3: LesserHeal,
}

var Iyuoid = Robot{
	id:     1,
	name:   "Iyuoid",
	health: 100,
	skill1: Destruction,
	skill2: Immobilize,
	skill3: GreaterHeal,
}

var Oza = Robot{
	id:     1,
	name:   "Oza",
	health: 100,
	skill1: Splint,
	skill2: Slice,
	skill3: Rampage,
}

var Skip = Robot{
	id:     1,
	name:   "Skip",
	health: 100,
	skill1: Destruction,
	skill2: Splint,
	skill3: LesserHeal,
}

var Boomer = Robot{
	id:     1,
	name:   "Boomer",
	health: 100,
	skill1: Purge,
	skill2: Splint,
	skill3: LesserHeal,
}

var Umx = Robot{
	id:     1,
	name:   "Umx",
	health: 100,
	skill1: LesserHeal,
	skill2: Strike,
	skill3: Immobilize,
}

var Mecha = Robot{
	id:     1,
	name:   "Mecha",
	health: 100,
	skill1: Rampage,
	skill2: Slam,
	skill3: LesserHeal,
}

var Anolator = Robot{
	id:     1,
	name:   "Anolator",
	health: 100,
	skill1: Destruction,
	skill2: Slam,
	skill3: GreaterHeal,
}

var Jet = Robot{
	id:     1,
	name:   "Jet",
	health: 100,
	skill1: Splint,
	skill2: Strike,
	skill3: LesserHeal,
}

var StarterRobots map[RobotId]Robot = map[RobotId]Robot{
	Jet.id:      Jet,
	Anolator.id: Anolator,
	Mecha.id:    Mecha,
	Umx.id:      Umx,
	Boomer.id:   Boomer,
	Skip.id:     Skip,
	Oza.id:      Oza,
	Iyuoid.id:   Iyuoid,
	Imotron.id:  Imotron,
	Tera.id:     Tera,
}

var StarterSkills [10]Skill = [10]Skill{Strike, Slam, Purge, LesserHeal, GreaterHeal, Splint, Destruction, Immobilize, Slice, Rampage}
