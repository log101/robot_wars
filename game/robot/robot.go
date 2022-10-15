package robot

import "strconv"

// ***********
// TYPE: SKILL
type Skill struct {
	name     string
	hpEffect int
}

// GETTERS
func (s *Skill) GetName() string {
	return s.name
}

func (s *Skill) GetHpEffect() int {
	return s.hpEffect
}

// SETTERS
func (s *Skill) SetName(name string) {
	s.name = name
}

func (s *Skill) SetHpEffect(effect int) {
	s.hpEffect = effect
}

// ***********
// TYPE: ROBOT
type RobotId int
type Robot struct {
	id     RobotId
	name   string
	health int
	skills [3]Skill
}

// METHODS
// Saldırı fonksiyonu bir yetenek ve hedef parametreleri alıyor
// yeteneğin gücünü rakip robotun canından çıkarıyor
func (r *Robot) Attact(skillIndex int, target *Robot) {
	effect := r.skills[skillIndex].hpEffect
	target.health += effect
}

// Her robotun bir id'si var bunu kullanarak da onlara ulaşabiliyoruz
func GetRobotById(r RobotId, robots map[RobotId]*Robot) *Robot {
	robot := robots[r]
	return robot
}

// Robot oluşturma fonksiyonu
func CreateRobot(id RobotId, name string, health int, skills [3]Skill) *Robot {
	newRobot := Robot{
		id:     id,
		name:   name,
		health: health,
		skills: skills,
	}

	return &newRobot
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

func (r *Robot) GetSkill(i int) Skill {
	return r.skills[i]
}

func (r *Robot) GetSkillFeatures(i int) string {
	return r.skills[i].name + " " + strconv.FormatInt(-int64(r.skills[i].hpEffect), 10)
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

func (r *Robot) SetSkill(i int, skill Skill) {
	r.skills[i] = skill
}
