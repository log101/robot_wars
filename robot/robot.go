package robot

// ***********
// TYPE: SKILL
type Skill struct {
	name       string
	hpEffect   int
	critChance int
}

// METHODS
func (s *Skill) Act(target *Robot) {
	target.health += s.hpEffect
}

// GETTERS
func (s *Skill) GetName() string {
	return s.name
}

func (s *Skill) GetHpEffect() int {
	return s.hpEffect
}

func (s *Skill) GetCritChance() int {
	return s.critChance
}

// SETTERS
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
type RobotId int
type Robot struct {
	id     RobotId
	name   string
	health int
	skill1 Skill
	skill2 Skill
	skill3 Skill
}

// GETTERS
func (r *Robot) GetId() RobotId {
	return r.id
}

func (r *Robot) GetName() string {
	return r.name
}

func (r *Robot) GetHealth() int {
	return r.health
}

func (r *Robot) GetSkill1() Skill {
	return r.skill1
}

func (r *Robot) GetSkill2() Skill {
	return r.skill2
}

func (r *Robot) GetSkill3() Skill {
	return r.skill3
}

// SETTERS
func (r *Robot) SetId(id RobotId) {
	r.id = id
}

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
