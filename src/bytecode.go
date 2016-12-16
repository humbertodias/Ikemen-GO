package main

import (
	"math"
	"unsafe"
)

type StateType int32

const (
	ST_S StateType = 1 << iota
	ST_C
	ST_A
	ST_L
	ST_N
	ST_U
	ST_D = ST_L
	ST_F = ST_N
	ST_P = ST_U
)

type AttackType int32

const (
	AT_NA AttackType = 1 << (iota + 6)
	AT_NT
	AT_NP
	AT_SA
	AT_ST
	AT_SP
	AT_HA
	AT_HT
	AT_HP
)

type MoveType int32

const (
	MT_I MoveType = 1 << (iota + 15)
	MT_H
	MT_A   = MT_I + 1
	MT_U   = MT_H + 1
	MT_MNS = MT_I
	MT_PLS = MT_H
)

type ValueType int

const (
	VT_Float ValueType = iota
	VT_Int
	VT_Bool
	VT_SFalse
)

type OpCode byte

const (
	OC_var OpCode = iota + 110
	OC_sysvar
	OC_fvar
	OC_sysfvar
	OC_int8
	OC_int
	OC_float
	OC_dup
	OC_swap
	OC_jmp8
	OC_jz8
	OC_jnz8
	OC_jmp
	OC_jz
	OC_jnz
	OC_eq
	OC_ne
	OC_gt
	OC_le
	OC_lt
	OC_ge
	OC_blnot
	OC_bland
	OC_blxor
	OC_blor
	OC_not
	OC_and
	OC_xor
	OC_or
	OC_shl
	OC_shr
	OC_add
	OC_sub
	OC_mul
	OC_div
	OC_mod
	OC_pow
	OC_abs
	OC_exp
	OC_ln
	OC_log
	OC_cos
	OC_sin
	OC_tan
	OC_acos
	OC_asin
	OC_atan
	OC_floor
	OC_ceil
	OC_ifelse
	OC_time
	OC_animtime
	OC_animelemtime
	OC_animelemno
	OC_statetype
	OC_movetype
	OC_ctrl
	OC_command
	OC_random
	OC_pos_x
	OC_pos_y
	OC_vel_x
	OC_vel_y
	OC_screenpos_x
	OC_screenpos_y
	OC_facing
	OC_anim
	OC_animexist
	OC_selfanimexist
	OC_alive
	OC_life
	OC_lifemax
	OC_power
	OC_powermax
	OC_canrecover
	OC_roundstate
	OC_ishelper
	OC_numhelper
	OC_numexplod
	OC_numprojid
	OC_numproj
	OC_teammode
	OC_teamside
	OC_hitdefattr
	OC_inguarddist
	OC_movecontact
	OC_movehit
	OC_moveguarded
	OC_movereversed
	OC_projcontacttime
	OC_projhittime
	OC_projguardedtime
	OC_projcanceltime
	OC_backedge
	OC_backedgedist
	OC_backedgebodydist
	OC_frontedge
	OC_frontedgedist
	OC_frontedgebodydist
	OC_leftedge
	OC_rightedge
	OC_topedge
	OC_bottomedge
	OC_camerapos_x
	OC_camerapos_y
	OC_camerazoom
	OC_gamewidth
	OC_gameheight
	OC_screenwidth
	OC_screenheight
	OC_stateno
	OC_prevstateno
	OC_id
	OC_playeridexist
	OC_gametime
	OC_numtarget
	OC_numenemy
	OC_numpartner
	OC_ailevel
	OC_palno
	OC_matchover
	OC_hitcount
	OC_uniqhitcount
	OC_hitpausetime
	OC_hitover
	OC_hitshakeover
	OC_hitfall
	OC_hitvel_x
	OC_hitvel_y
	OC_roundno
	OC_roundsexisted
	OC_matchno
	OC_ishometeam
	OC_parent
	OC_root
	OC_helper
	OC_target
	OC_partner
	OC_enemy
	OC_enemynear
	OC_playerid
	OC_p2
	OC_const_
	OC_gethitvar_
	OC_stagevar_
	OC_ex_
	OC_var0     OpCode = 0
	OC_sysvar0  OpCode = 60
	OC_fvar0    OpCode = 65
	OC_sysfvar0 OpCode = 105
)
const (
	OC_const_data_life OpCode = iota
	OC_const_data_power
	OC_const_data_attack
	OC_const_data_defence
	OC_const_data_fall_defence_mul
	OC_const_data_liedown_time
	OC_const_data_airjuggle
	OC_const_data_sparkno
	OC_const_data_guard_sparkno
	OC_const_data_ko_echo
	OC_const_data_intpersistindex
	OC_const_data_floatpersistindex
	OC_const_size_xscale
	OC_const_size_yscale
	OC_const_size_ground_back
	OC_const_size_ground_front
	OC_const_size_air_back
	OC_const_size_air_front
	OC_const_size_z_width
	OC_const_size_height
	OC_const_size_attack_dist
	OC_const_size_attack_z_width_back
	OC_const_size_attack_z_width_front
	OC_const_size_proj_attack_dist
	OC_const_size_proj_doscale
	OC_const_size_head_pos_x
	OC_const_size_head_pos_y
	OC_const_size_mid_pos_x
	OC_const_size_mid_pos_y
	OC_const_size_shadowoffset
	OC_const_size_draw_offset_x
	OC_const_size_draw_offset_y
	OC_const_velocity_walk_fwd_x
	OC_const_velocity_walk_back_x
	OC_const_velocity_walk_up_x
	OC_const_velocity_walk_down_x
	OC_const_velocity_run_fwd_x
	OC_const_velocity_run_fwd_y
	OC_const_velocity_run_back_x
	OC_const_velocity_run_back_y
	OC_const_velocity_run_up_x
	OC_const_velocity_run_up_y
	OC_const_velocity_run_down_x
	OC_const_velocity_run_down_y
	OC_const_velocity_jump_y
	OC_const_velocity_jump_neu_x
	OC_const_velocity_jump_back_x
	OC_const_velocity_jump_fwd_x
	OC_const_velocity_jump_up_x
	OC_const_velocity_jump_down_x
	OC_const_velocity_runjump_back_x
	OC_const_velocity_runjump_back_y
	OC_const_velocity_runjump_fwd_x
	OC_const_velocity_runjump_fwd_y
	OC_const_velocity_runjump_up_x
	OC_const_velocity_runjump_down_x
	OC_const_velocity_airjump_y
	OC_const_velocity_airjump_neu_x
	OC_const_velocity_airjump_back_x
	OC_const_velocity_airjump_fwd_x
	OC_const_velocity_airjump_up_x
	OC_const_velocity_airjump_down_x
	OC_const_velocity_air_gethit_groundrecover_x
	OC_const_velocity_air_gethit_groundrecover_y
	OC_const_velocity_air_gethit_airrecover_mul_x
	OC_const_velocity_air_gethit_airrecover_mul_y
	OC_const_velocity_air_gethit_airrecover_add_x
	OC_const_velocity_air_gethit_airrecover_add_y
	OC_const_velocity_air_gethit_airrecover_back
	OC_const_velocity_air_gethit_airrecover_fwd
	OC_const_velocity_air_gethit_airrecover_up
	OC_const_velocity_air_gethit_airrecover_down
	OC_const_movement_airjump_num
	OC_const_movement_airjump_height
	OC_const_movement_yaccel
	OC_const_movement_stand_friction
	OC_const_movement_crouch_friction
	OC_const_movement_stand_friction_threshold
	OC_const_movement_crouch_friction_threshold
	OC_const_movement_jump_changeanim_threshold
	OC_const_movement_air_gethit_groundlevel
	OC_const_movement_air_gethit_groundrecover_ground_threshold
	OC_const_movement_air_gethit_groundrecover_groundlevel
	OC_const_movement_air_gethit_airrecover_threshold
	OC_const_movement_air_gethit_airrecover_yaccel
	OC_const_movement_air_gethit_trip_groundlevel
	OC_const_movement_down_bounce_offset_x
	OC_const_movement_down_bounce_offset_y
	OC_const_movement_down_bounce_yaccel
	OC_const_movement_down_bounce_groundlevel
	OC_const_movement_down_friction_threshold
)
const (
	OC_gethitvar_animtype OpCode = iota
	OC_gethitvar_airtype
	OC_gethitvar_groundtype
	OC_gethitvar_damage
	OC_gethitvar_hitcount
	OC_gethitvar_fallcount
	OC_gethitvar_hitshaketime
	OC_gethitvar_hittime
	OC_gethitvar_slidetime
	OC_gethitvar_ctrltime
	OC_gethitvar_recovertime
	OC_gethitvar_xoff
	OC_gethitvar_yoff
	OC_gethitvar_xvel
	OC_gethitvar_yvel
	OC_gethitvar_yaccel
	OC_gethitvar_chainid
	OC_gethitvar_guarded
	OC_gethitvar_isbound
	OC_gethitvar_fall
	OC_gethitvar_fall_damage
	OC_gethitvar_fall_xvel
	OC_gethitvar_fall_yvel
	OC_gethitvar_fall_recover
	OC_gethitvar_fall_recovertime
	OC_gethitvar_fall_kill
	OC_gethitvar_fall_envshake_time
	OC_gethitvar_fall_envshake_freq
	OC_gethitvar_fall_envshake_ampl
	OC_gethitvar_fall_envshake_phase
)
const (
	OC_stagevar_info_author OpCode = iota
	OC_stagevar_info_displayname
	OC_stagevar_info_name
)
const (
	OC_ex_name OpCode = iota
	OC_ex_authorname
	OC_ex_p2name
	OC_ex_p3name
	OC_ex_p4name
	OC_ex_p2dist_x
	OC_ex_p2dist_y
	OC_ex_p2bodydist_x
	OC_ex_p2bodydist_y
	OC_ex_parentdist_x
	OC_ex_parentdist_y
	OC_ex_rootdist_x
	OC_ex_rootdist_y
	OC_ex_win
	OC_ex_winko
	OC_ex_wintime
	OC_ex_winperfect
	OC_ex_lose
	OC_ex_loseko
	OC_ex_losetime
	OC_ex_drawgame
	OC_ex_tickspersecond
)

type StringPool struct {
	List []string
	Map  map[string]int
}

func NewStringPool() *StringPool {
	return &StringPool{Map: make(map[string]int)}
}
func (sp *StringPool) Clear(s string) {
	sp.List, sp.Map = nil, make(map[string]int)
}
func (sp *StringPool) Add(s string) int {
	i, ok := sp.Map[s]
	if !ok {
		i = len(sp.List)
		sp.List = append(sp.List, s)
		sp.Map[s] = i
	}
	return i
}

type BytecodeValue struct {
	t ValueType
	v float64
}

func (bv BytecodeValue) IsSF() bool { return bv.t == VT_SFalse }
func (bv BytecodeValue) ToF() float32 {
	if bv.IsSF() {
		return 0
	}
	return float32(bv.v)
}
func (bv BytecodeValue) ToI() int32 {
	if bv.IsSF() {
		return 0
	}
	return int32(bv.v)
}
func (bv BytecodeValue) ToB() bool {
	if bv.IsSF() || bv.v == 0 {
		return false
	}
	return true
}
func (bv *BytecodeValue) SetF(f float32) {
	*bv = BytecodeValue{VT_Float, float64(f)}
}
func (bv *BytecodeValue) SetI(i int32) {
	*bv = BytecodeValue{VT_Int, float64(i)}
}
func (bv *BytecodeValue) SetB(b bool) {
	bv.t = VT_Bool
	if b {
		bv.v = 1
	} else {
		bv.v = 0
	}
}

func BytecodeSF() BytecodeValue {
	return BytecodeValue{VT_SFalse, math.NaN()}
}

type BytecodeStack []BytecodeValue

func (bs *BytecodeStack) Clear()                { *bs = (*bs)[:0] }
func (bs *BytecodeStack) Push(bv BytecodeValue) { *bs = append(*bs, bv) }
func (bs BytecodeStack) Top() *BytecodeValue {
	return &bs[len(bs)-1]
}
func (bs *BytecodeStack) Pop() (bv BytecodeValue) {
	bv, *bs = *bs.Top(), (*bs)[:len(*bs)-1]
	return
}
func (bs *BytecodeStack) Dup() {
	bs.Push(*bs.Top())
}
func (bs *BytecodeStack) Swap() {
	*bs.Top(), (*bs)[len(*bs)-2] = (*bs)[len(*bs)-2], *bs.Top()
}

type BytecodeExp []OpCode

func (be *BytecodeExp) append(op ...OpCode) {
	*be = append(*be, op...)
}
func (be *BytecodeExp) appendFloat(f float32) {
	be.append((*(*[4]OpCode)(unsafe.Pointer(&f)))[:]...)
}
func (be *BytecodeExp) appendInt(i int32) {
	be.append((*(*[4]OpCode)(unsafe.Pointer(&i)))[:]...)
}
func (be BytecodeExp) toF() float32 {
	return *(*float32)(unsafe.Pointer(&be[0]))
}
func (be BytecodeExp) toI() int32 {
	return *(*int32)(unsafe.Pointer(&be[0]))
}
func (be *BytecodeExp) appendValue(bv BytecodeValue) (ok bool) {
	switch bv.t {
	case VT_Float:
		be.append(OC_float)
		be.appendFloat(float32(bv.v))
	case VT_Int:
		if bv.v >= -128 || bv.v <= 127 {
			be.append(OC_int8, OpCode(bv.v))
		} else {
			be.append(OC_int)
			be.appendInt(int32(bv.v))
		}
	case VT_Bool:
		if bv.v != 0 {
			be.append(OC_int8, 1)
		} else {
			be.append(OC_int8, 0)
		}
	default:
		return false
	}
	return true
}
func (_ BytecodeExp) blnot(v *BytecodeValue) {
	if v.ToB() {
		v.v = 0
	} else {
		v.v = 1
	}
	v.t = VT_Int
}
func (_ BytecodeExp) pow(v1 *BytecodeValue, v2 BytecodeValue, pn int) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetF(float32(math.Pow(float64(v1.ToF()), float64(v2.ToF()))))
	} else if v2.ToF() < 0 {
		if sys.cgi[pn].ver[0] == 1 {
			v1.SetF(float32(math.Pow(float64(v1.ToI()), float64(v2.ToI()))))
		} else {
			f := float32(math.Pow(float64(v1.ToI()), float64(v2.ToI())))
			v1.SetI(*(*int32)(unsafe.Pointer(&f)) << 29)
		}
	} else {
		i1, i2, hb := v1.ToI(), v2.ToI(), int32(-1)
		for uint32(i2)>>uint(hb+1) != 0 {
			hb++
		}
		var i, bit, tmp int32 = 1, 0, i1
		for ; bit <= hb; bit++ {
			var shift uint
			if bit == hb || sys.cgi[pn].ver[0] == 1 {
				shift = uint(bit)
			} else {
				shift = uint((hb - 1) - bit)
			}
			if i2&(1<<shift) != 0 {
				i *= tmp
			}
			tmp *= tmp
		}
		v1.SetI(i)
	}
}
func (_ BytecodeExp) mul(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetF(v1.ToF() * v2.ToF())
	} else {
		v1.SetI(v1.ToI() * v2.ToI())
	}
}
func (_ BytecodeExp) div(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetF(v1.ToF() / v2.ToF())
	} else if v2.ToI() == 0 {
		*v1 = BytecodeSF()
	} else {
		v1.SetI(v1.ToI() / v2.ToI())
	}
}
func (_ BytecodeExp) mod(v1 *BytecodeValue, v2 BytecodeValue) {
	if v2.ToI() == 0 {
		*v1 = BytecodeSF()
	} else {
		v1.SetI(v1.ToI() % v2.ToI())
	}
}
func (_ BytecodeExp) add(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetF(v1.ToF() + v2.ToF())
	} else {
		v1.SetI(v1.ToI() + v2.ToI())
	}
}
func (_ BytecodeExp) sub(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetF(v1.ToF() - v2.ToF())
	} else {
		v1.SetI(v1.ToI() - v2.ToI())
	}
}
func (_ BytecodeExp) gt(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() > v2.ToF())
	} else {
		v1.SetB(v1.ToI() > v2.ToI())
	}
}
func (_ BytecodeExp) ge(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() >= v2.ToF())
	} else {
		v1.SetB(v1.ToI() >= v2.ToI())
	}
}
func (_ BytecodeExp) lt(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() < v2.ToF())
	} else {
		v1.SetB(v1.ToI() < v2.ToI())
	}
}
func (_ BytecodeExp) le(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() <= v2.ToF())
	} else {
		v1.SetB(v1.ToI() <= v2.ToI())
	}
}
func (_ BytecodeExp) eq(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() == v2.ToF())
	} else {
		v1.SetB(v1.ToI() == v2.ToI())
	}
}
func (_ BytecodeExp) ne(v1 *BytecodeValue, v2 BytecodeValue) {
	if ValueType(Min(int32(v1.t), int32(v2.t))) == VT_Float {
		v1.SetB(v1.ToF() != v2.ToF())
	} else {
		v1.SetB(v1.ToI() != v2.ToI())
	}
}
func (_ BytecodeExp) and(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetI(v1.ToI() & v2.ToI())
}
func (_ BytecodeExp) xor(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetI(v1.ToI() ^ v2.ToI())
}
func (_ BytecodeExp) or(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetI(v1.ToI() | v2.ToI())
}
func (_ BytecodeExp) bland(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetB(v1.ToB() && v2.ToB())
}
func (_ BytecodeExp) blxor(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetB(v1.ToB() != v2.ToB())
}
func (_ BytecodeExp) blor(v1 *BytecodeValue, v2 BytecodeValue) {
	v1.SetB(v1.ToB() || v2.ToB())
}
func (be BytecodeExp) run(c *Char, scpn int) BytecodeValue {
	sys.bcStack.Clear()
	for i := 1; i <= len(be); i++ {
		switch be[i-1] {
		case OC_int8:
			sys.bcStack.Push(BytecodeValue{VT_Int, float64(int8(be[i]))})
			i++
		case OC_int:
			sys.bcStack.Push(BytecodeValue{VT_Int, float64(be[i:].toI())})
			i += 4
		case OC_float:
			sys.bcStack.Push(BytecodeValue{VT_Float, float64(be[i:].toF())})
			i += 4
		case OC_blnot:
			be.blnot(sys.bcStack.Top())
		case OC_pow:
			v2 := sys.bcStack.Pop()
			be.pow(sys.bcStack.Top(), v2, scpn)
		case OC_mul:
			v2 := sys.bcStack.Pop()
			be.mul(sys.bcStack.Top(), v2)
		case OC_div:
			v2 := sys.bcStack.Pop()
			be.div(sys.bcStack.Top(), v2)
		case OC_mod:
			v2 := sys.bcStack.Pop()
			be.mod(sys.bcStack.Top(), v2)
		case OC_add:
			v2 := sys.bcStack.Pop()
			be.add(sys.bcStack.Top(), v2)
		case OC_sub:
			v2 := sys.bcStack.Pop()
			be.sub(sys.bcStack.Top(), v2)
		case OC_gt:
			v2 := sys.bcStack.Pop()
			be.gt(sys.bcStack.Top(), v2)
		case OC_ge:
			v2 := sys.bcStack.Pop()
			be.ge(sys.bcStack.Top(), v2)
		case OC_lt:
			v2 := sys.bcStack.Pop()
			be.lt(sys.bcStack.Top(), v2)
		case OC_le:
			v2 := sys.bcStack.Pop()
			be.le(sys.bcStack.Top(), v2)
		case OC_eq:
			v2 := sys.bcStack.Pop()
			be.eq(sys.bcStack.Top(), v2)
		case OC_ne:
			v2 := sys.bcStack.Pop()
			be.ne(sys.bcStack.Top(), v2)
		case OC_and:
			v2 := sys.bcStack.Pop()
			be.and(sys.bcStack.Top(), v2)
		case OC_xor:
			v2 := sys.bcStack.Pop()
			be.xor(sys.bcStack.Top(), v2)
		case OC_or:
			v2 := sys.bcStack.Pop()
			be.or(sys.bcStack.Top(), v2)
		case OC_bland:
			v2 := sys.bcStack.Pop()
			be.bland(sys.bcStack.Top(), v2)
		case OC_blxor:
			v2 := sys.bcStack.Pop()
			be.blxor(sys.bcStack.Top(), v2)
		case OC_blor:
			v2 := sys.bcStack.Pop()
			be.blor(sys.bcStack.Top(), v2)
		case OC_dup:
			sys.bcStack.Dup()
		case OC_swap:
			sys.bcStack.Swap()
		default:
			unimplemented()
		}
	}
	return sys.bcStack.Pop()
}
func (be BytecodeExp) evalF(c *Char, scpn int) float32 {
	return be.run(c, scpn).ToF()
}
func (be BytecodeExp) evalI(c *Char, scpn int) int32 {
	return be.run(c, scpn).ToI()
}
func (be BytecodeExp) evalB(c *Char, scpn int) bool {
	return be.run(c, scpn).ToB()
}

type StateController interface {
	Run(c *Char, scpn int) (changeState bool)
}

const (
	SCID_trigger byte = 0
	SCID_const   byte = 128
)

type StateControllerBase struct {
	playerNo       int
	persistent     int32
	ignorehitpause bool
	code           []byte
}

func newStateControllerBase(pn int) *StateControllerBase {
	return &StateControllerBase{playerNo: pn, persistent: 1}
}
func (scb StateControllerBase) beToExp(be ...BytecodeExp) []BytecodeExp {
	return be
}
func (scb StateControllerBase) fToExp(f ...float32) (exp []BytecodeExp) {
	for _, v := range f {
		var be BytecodeExp
		be.appendFloat(v)
		exp = append(exp, be)
	}
	return
}
func (scb StateControllerBase) iToExp(i ...int32) (exp []BytecodeExp) {
	for _, v := range i {
		var be BytecodeExp
		be.appendInt(v)
		exp = append(exp, be)
	}
	return
}
func (scb *StateControllerBase) add(id byte, exp []BytecodeExp) {
	scb.code = append(scb.code, id, byte(len(exp)))
	for _, e := range exp {
		l := int32(len(e))
		scb.code = append(scb.code, (*(*[4]byte)(unsafe.Pointer(&l)))[:]...)
		scb.code = append(scb.code, (*(*[]byte)(unsafe.Pointer(&e)))...)
	}
}
func (scb StateControllerBase) run(f func(byte, []BytecodeExp) bool) bool {
	for i := 0; i < len(scb.code); {
		id := scb.code[i]
		i++
		n := scb.code[i]
		i++
		exp := make([]BytecodeExp, n)
		for m := 0; m < int(n); m++ {
			l := *(*int32)(unsafe.Pointer(&scb.code[i]))
			i += 4
			exp[m] = (*(*BytecodeExp)(unsafe.Pointer(&scb.code)))[i : i+int(l)]
			i += int(l)
		}
		if !f(id, exp) {
			return false
		}
	}
	return true
}

type stateDef StateControllerBase

const (
	stateDef_hitcountpersist byte = iota + 1
	stateDef_movehitpersist
	stateDef_hitdefpersist
	stateDef_sprpriority
	stateDef_facep2
	stateDef_juggle
	stateDef_velset
	stateDef_anim
	stateDef_ctrl
	stateDef_poweradd
	stateDef_hitcountpersist_c = stateDef_hitcountpersist + SCID_const
	stateDef_movehitpersist_c  = stateDef_movehitpersist + SCID_const
	stateDef_hitdefpersist_c   = stateDef_hitdefpersist + SCID_const
	stateDef_sprpriority_c     = stateDef_sprpriority + SCID_const
	stateDef_facep2_c          = stateDef_facep2 + SCID_const
	stateDef_juggle_c          = stateDef_juggle + SCID_const
	stateDef_velset_c          = stateDef_velset + SCID_const
	stateDef_anim_c            = stateDef_anim + SCID_const
	stateDef_ctrl_c            = stateDef_ctrl + SCID_const
	stateDef_poweradd_c        = stateDef_poweradd + SCID_const
)

func (sd stateDef) Run(c *Char, scpn int) bool {
	StateControllerBase(sd).run(func(id byte, exp []BytecodeExp) bool {
		switch id {
		case stateDef_hitcountpersist, stateDef_hitcountpersist_c:
			if id == stateDef_hitcountpersist_c || !exp[0].evalB(c, scpn) {
				c.clearHitCount()
			}
		case stateDef_movehitpersist, stateDef_movehitpersist_c:
			if id == stateDef_movehitpersist_c || !exp[0].evalB(c, scpn) {
				c.clearMoveHit()
			}
		case stateDef_hitdefpersist, stateDef_hitdefpersist_c:
			if id == stateDef_hitdefpersist_c || !exp[0].evalB(c, scpn) {
				c.clearHitDef()
			}
		case stateDef_sprpriority:
			c.setSprPriority(exp[0].evalI(c, scpn))
		case stateDef_sprpriority_c:
			c.setSprPriority(exp[0].toI())
		case stateDef_facep2, stateDef_facep2_c:
			if id == stateDef_facep2_c || exp[0].evalB(c, scpn) {
				c.faceP2()
			}
		case stateDef_juggle:
			c.setJuggle(exp[0].evalI(c, scpn))
		case stateDef_juggle_c:
			c.setJuggle(exp[0].toI())
		case stateDef_velset:
			c.setXV(exp[0].evalF(c, scpn))
			if len(exp) > 1 {
				c.setYV(exp[1].evalF(c, scpn))
				if len(exp) > 2 {
					exp[2].run(c, scpn)
				}
			}
		case stateDef_velset_c:
			c.setXV(exp[0].toF())
			if len(exp) > 1 {
				c.setYV(exp[1].toF())
			}
		case stateDef_anim:
			c.changeAnim(exp[0].evalI(c, scpn))
		case stateDef_anim_c:
			c.changeAnim(exp[0].toI())
		case stateDef_ctrl:
			c.setCtrl(exp[0].evalB(c, scpn))
		case stateDef_ctrl_c:
			c.setCtrl(exp[0].toI() != 0)
		case stateDef_poweradd:
			c.addPower(exp[0].evalI(c, scpn))
		case stateDef_poweradd_c:
			c.addPower(exp[0].toI())
		}
		return true
	})
	return false
}

type hitBy StateControllerBase

const (
	_hitBy_value byte = iota
	_hitBy_value2
	hitBy_time
	hitBy_value_c  = _hitBy_value + SCID_const
	hitBy_value2_c = _hitBy_value2 + SCID_const
	hitBy_time_c   = hitBy_time + SCID_const
)

func (nhb hitBy) Run(c *Char, scpn int) bool {
	unimplemented()
	return false
}

type notHitBy hitBy

func (nhb notHitBy) Run(c *Char, scpn int) bool {
	unimplemented()
	return false
}

type StateBytecode struct {
	stateType StateType
	moveType  MoveType
	physics   StateType
	stateDef  StateController
	ctrls     []StateController
}

func newStateBytecode() *StateBytecode {
	return &StateBytecode{stateType: ST_S, moveType: MT_I, physics: ST_N}
}

type Bytecode struct{ states map[int32]StateBytecode }

func newBytecode() *Bytecode {
	return &Bytecode{states: make(map[int32]StateBytecode)}
}
