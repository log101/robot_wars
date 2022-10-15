package robot

// SEED DATA
// SKILLS
var Strike = Skill{
	name:     "strike",
	hpEffect: -50,
}

var Slam = Skill{
	name:     "slam",
	hpEffect: -20,
}

var Purge = Skill{
	name:     "purge",
	hpEffect: -30,
}

var LesserFire = Skill{
	name:     "lesser fire",
	hpEffect: -20,
}

var GreaterFire = Skill{
	name:     "greater fire",
	hpEffect: -50,
}

var Splint = Skill{
	name:     "splint",
	hpEffect: -45,
}

var Destruction = Skill{
	name:     "destruction",
	hpEffect: -70,
}

var Immobilize = Skill{
	name:     "immobilize",
	hpEffect: -10,
}

var Slice = Skill{
	name:     "slice",
	hpEffect: -5,
}

var Rampage = Skill{
	name:     "rampage",
	hpEffect: -90,
}

// ROBOTS
var Tera = Robot{
	id:     1,
	name:   "Tera",
	health: 100,
	skills: [3]Skill{Rampage, Slice, Immobilize},
}

var Imotron = Robot{
	id:     2,
	name:   "Imotron",
	health: 100,
	skills: [3]Skill{Strike, Slam, LesserFire},
}

var Iyuoid = Robot{
	id:     3,
	name:   "Iyuoid",
	health: 100,
	skills: [3]Skill{Destruction, Immobilize, GreaterFire},
}

var Oza = Robot{
	id:     4,
	name:   "Oza",
	health: 100,
	skills: [3]Skill{Splint, Slice, Rampage},
}

var Skip = Robot{
	id:     5,
	name:   "Skip",
	health: 100,
	skills: [3]Skill{Destruction, Splint, LesserFire},
}

var Boomer = Robot{
	id:     6,
	name:   "Boomer",
	health: 100,
	skills: [3]Skill{Purge, Splint, LesserFire},
}

var Umx = Robot{
	id:     7,
	name:   "Umx",
	health: 100,
	skills: [3]Skill{LesserFire, Strike, Immobilize},
}

var Mecha = Robot{
	id:     8,
	name:   "Mecha",
	health: 100,
	skills: [3]Skill{Rampage, Slam, LesserFire},
}

var Anolator = Robot{
	id:     9,
	name:   "Anolator",
	health: 100,
	skills: [3]Skill{Destruction, Slam, GreaterFire},
}

var Jet = Robot{
	id:     10,
	name:   "Jet",
	health: 100,
	skills: [3]Skill{Splint, Strike, LesserFire},
}

var StarterRobots map[RobotId]*Robot = map[RobotId]*Robot{
	Jet.id:      &Jet,
	Anolator.id: &Anolator,
	Mecha.id:    &Mecha,
	Umx.id:      &Umx,
	Boomer.id:   &Boomer,
	Skip.id:     &Skip,
	Oza.id:      &Oza,
	Iyuoid.id:   &Iyuoid,
	Imotron.id:  &Imotron,
	Tera.id:     &Tera,
}

var StarterRobotsSample map[RobotId]*Robot = map[RobotId]*Robot{
	Oza.id:     &Oza,
	Iyuoid.id:  &Iyuoid,
	Imotron.id: &Imotron,
	Tera.id:    &Tera,
}

var StarterSkills [10]Skill = [10]Skill{Strike, Slam, Purge, LesserFire, GreaterFire, Splint, Destruction, Immobilize, Slice, Rampage}
