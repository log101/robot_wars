package robot

// ***********
// TYPE: SKILL
type Skill struct {
	name       string
	hpEffect   int
	critChance int
}

// METHODS
func (r *Robot) Act(skill Skill, target *Robot) {
	target.health += skill.hpEffect
}

// GETTERS & SETTERS
func (s *Skill) SetName(name string) {
	s.name = name
}

func (s *Skill) SetHpEffect(effect int) {
	s.hpEffect = effect
}

func (s *Skill) SetCritChance(chance int) {
	s.critChance = chance
}

// ***********
// TYPE: ROBOT
type Robot struct {
	name   string
	health int
	skill1 Skill
	skill2 Skill
	skill3 Skill
}

// GETTERS & SETTERS
func (r *Robot) SetName(name string) {
	r.name = name
}

func (r *Robot) SetHealth(health int) {
	r.health = health
}

func (r *Robot) SetSkill1(skill Skill) {
	r.skill1 = skill
}

func (r *Robot) SetSkill2(skill Skill) {
	r.skill2 = skill
}

func (r *Robot) SetSkill3(skill Skill) {
	r.skill3 = skill
}
