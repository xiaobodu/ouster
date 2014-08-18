package main

import (
    "time"
)

type WearPart uint

const (
    WEAR_CIRCLET = iota
    WEAR_COAT
    WEAR_LEFTHAND
    WEAR_RIGHTHAND
    WEAR_BOOTS
    WEAR_ARMSBAND1
    WEAR_ARMSBAND2
    WEAR_RING1
    WEAR_RING2
    WEAR_NECKLACE1
    WEAR_NECKLACE2
    WEAR_NECKLACE3
    WEAR_STONE1
    WEAR_STONE2
    WEAR_STONE3
    WEAR_STONE4
    WEAR_ZAP1
    WEAR_ZAP2
    WEAR_ZAP3
    WEAR_ZAP4
    WEAR_FASCIA
    WEAR_MITTEN
    OUSTERS_WEAR_MAX
)

type Ouster struct {
    PlayerCreature //继承自PlayerCreature
    
    Name            string
    Competence      byte
    CompetenceShape byte

    HairColor Color_t
    Alignment Alignment_t

    STR [3]Attr_t
    DEX [3]Attr_t
    INT [3]Attr_t

    HP  [3]HP_t
    MP  [3]MP_t

    Damage        [3]Damage_t
    ToHit         [2]ToHit_t
    Defense       [2]Defense_t
    Protection    [2]Protection_t
    AttackSpeed   [2]Speed_t
    CriticalRatio [2]int

    GoalExp Exp_t
    Level   Level_t

    Bonus      Bonus_t
    SkillBonus SkillBonus_t

    Gold Gold_t
    Fame Fame_t

    VisionWidth  ZoneCoord_t
    VisionHeight ZoneCoord_t

    // 技能糟
    SkillSlot map[SkillType_t]*OusterSkillSlot

    WearItem [OUSTERS_WEAR_MAX]*Item

    SilverDamage Silver_t

    HPStealAmount Steal_t
    HPStealRatio  Steal_t

    MPStealAmount Steal_t
    MPStealRatio  Steal_t

    HPRegen Regen_t
    MPRegen Regen_t

    Luck Luck_t

    ElementalFire  Elemental_t
    ElementalWater Elemental_t
    ElementalEarth Elemental_t
    ElementalWind  Elemental_t

    FireDamage  Damage_t
    WaterDamage Damage_t
    EarthDamage Damage_t

    SilverResist Resist_t

    PassiveSkillMap    map[SkillType_t]struct{}
    PassiveRatio       int
    ExpSaveCount       uint16
    FameSaveCount      uint16
    AlignmentSaveCount uint16

    MPRegenTime time.Time
}

func (ouster Ouster) CreatureClass() CreatureClass {
	return CREATURE_CLASS_OUSTER
}