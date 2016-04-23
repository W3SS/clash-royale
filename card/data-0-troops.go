package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/rng"
	"github.com/asukakenji/clash-royale/speed"
	"github.com/asukakenji/clash-royale/targets"
	"github.com/asukakenji/clash-royale/typ"
)

var troops = [...]card{
	// --- Common Troops ---
	card{
		attr.Name:    "Knight",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.`,
		attr.Elixir:  3,
		attr.BaseHP:  600,
		attr.DPS:     []interface{}{68, 74, 81, 90, 99, 109, 120, 130, 144, 158, 174, 190},
		attr.BaseDam: 75,
		attr.HSpeed:  1.1,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Medium,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:    "Archers",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `A pair of unarmored ranged attackers. They'll help you with ground and air unit attacks, but you're on your own with coloring your hair.`,
		attr.Elixir:  3,
		attr.BaseHP:  125,
		attr.DPS:     []interface{}{33, 36, 40, 44, 48, 53, 58, 64, 70, 77, 85, 93},
		attr.BaseDam: 40,
		attr.HSpeed:  1.2,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.Medium,
		attr.Range:   5.5,
		attr.DTime:   1,
		attr.Count:   2,
	}.init(),
	card{
		attr.Name:     "Bomber",
		attr.Arena:    arena.Arena0,
		attr.Rarity:   rarity.Common,
		attr.Type:     typ.Troop,
		attr.Desc:     `Small, lightly protected skeleton that throws bombs. Deals area damage that can wipe out a swarm of enemies.`,
		attr.Elixir:   3,
		attr.BaseHP:   150,
		attr.DPS:      []interface{}{52, 57, 63, 70, 76, 84, 92, 101, 111, 122, 134, 147},
		attr.BaseADam: 100,
		attr.HSpeed:   1.9,
		attr.Targets:  targets.Ground,
		attr.Speed:    speed.Medium,
		attr.Range:    5,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:    "Goblins",
		attr.Arena:   arena.Arena1,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Three fast, unarmored melee attackers. Small, fast, green and mean!`,
		attr.Elixir:  2,
		attr.BaseHP:  80,
		attr.DPS:     []interface{}{45, 50, 54, 60, 66, 72, 80, 87, 96, 105, 116, 127},
		attr.BaseDam: 50,
		attr.HSpeed:  1.1,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.VeryFast,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
		attr.Count:   3,
	}.init(),
	card{
		attr.Name:    "Spear Goblins",
		attr.Arena:   arena.Arena1,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Three unarmored ranged attackers. Who the heck taught these guys to throw spears!?! Who thought that was a good idea?!`,
		attr.Elixir:  2,
		attr.BaseHP:  52,
		attr.DPS:     []interface{}{18, 20, 22, 23, 26, 29, 32, 35, 38, 42, 46, 51},
		attr.BaseDam: 24,
		attr.HSpeed:  1.3,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.VeryFast,
		attr.Range:   5.5,
		attr.DTime:   1,
		attr.Count:   3,
	}.init(),
	card{
		attr.Name:    "Skeletons",
		attr.Arena:   arena.Arena2,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Four fast, very weak melee fighters. Swarm your enemies with this pile of bones!`,
		attr.Elixir:  1,
		attr.BaseHP:  30,
		attr.DPS:     []interface{}{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		attr.BaseDam: 30,
		attr.HSpeed:  1,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Fast,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
		attr.Count:   4,
	}.init(),
	card{
		attr.Name:    "Minions",
		attr.Arena:   arena.Arena2,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Three fast, unarmored flying attackers. Roses are red, minions are blue, they can fly, and will crush you!`,
		attr.Elixir:  3,
		attr.BaseHP:  90,
		attr.DPS:     []interface{}{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		attr.BaseDam: 40,
		attr.HSpeed:  1,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.Fast,
		attr.Range:   2.5,
		attr.DTime:   1,
		attr.Count:   3,
	}.init(),
	card{
		attr.Name:    "Barbarians",
		attr.Arena:   arena.Arena3,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `A horde of melee attackers with mean mustaches and even meaner tempers.`,
		attr.Elixir:  5,
		attr.BaseHP:  300,
		attr.DPS:     []interface{}{50, 54, 60, 66, 72, 80, 88, 96, 106, 116, 128, 140},
		attr.BaseDam: 75,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Medium,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
		attr.Count:   4,
	}.init(),
	card{
		attr.Name:    "Minion Horde",
		attr.Arena:   arena.Arena4,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Six fast, unarmored flying attackers. Three's a crowd, six is a horde!`,
		attr.Elixir:  5,
		attr.BaseHP:  90,
		attr.DPS:     []interface{}{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		attr.BaseDam: 40,
		attr.HSpeed:  1,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.Fast,
		attr.Range:   2.5,
		attr.DTime:   1,
		attr.Count:   6,
	}.init(),
	card{
		attr.Name:    "Royale Giant",
		attr.Arena:   arena.Arena7,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Troop,
		attr.Desc:    `Sighting his massive cannon at enemy buildings, the Royal Giant comes in like a wrecking ball.`,
		attr.Elixir:  6,
		attr.BaseHP:  1200,
		attr.DPS:     []interface{}{52, 56, 62, 68, 75, 83, 91, 100, 110, 122, 133, 146},
		attr.BaseDam: 78,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Buildings,
		attr.Speed:   speed.Slow,
		attr.Range:   6,
		attr.DTime:   1,
	}.init(),

	// --- Rare Troops ---
	card{
		attr.Name:    "Giant",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `Slow but durable, only attacks buildings. A real one-man wrecking crew!`,
		attr.Elixir:  5,
		attr.BaseHP:  2000,
		attr.DPS:     []interface{}{84, 92, 101, 111, 122, 134, 147, 162, 177, 195},
		attr.BaseDam: 126,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Buildings,
		attr.Speed:   speed.Slow,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:    "Musketeer",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `Don't be fooled by her delicately coiffed hair, the Musketeer is a mean shot with her trusty boomstick.`,
		attr.Elixir:  4,
		attr.BaseHP:  340,
		attr.DPS:     []interface{}{90, 100, 110, 120, 132, 145, 160, 175, 192, 211},
		attr.BaseDam: 100,
		attr.HSpeed:  1.1,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.Medium,
		attr.Range:   6.5,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:    "Mini P.E.K.K.A",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `The arena is a certified butterfly-free zone. No distractions for P.E.K.K.A, only destruction.`,
		attr.Elixir:  4,
		attr.BaseHP:  600,
		attr.DPS:     []interface{}{180, 198, 218, 240, 263, 288, 317, 348, 382, 420},
		attr.BaseDam: 325,
		attr.HSpeed:  1.8,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Fast,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:    "Valkyrie",
		attr.Arena:   arena.Arena1,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `Tough melee fighter, deals area damage around her. Swarm or horde, no problem! She can take them all out with a few spins.`,
		attr.Elixir:  4,
		attr.BaseHP:  880,
		attr.DPS:     []interface{}{80, 88, 96, 106, 116, 128, 140, 154, 169, 186},
		attr.BaseDam: 120,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Medium,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:    "Hog Rider",
		attr.Arena:   arena.Arena4,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `Fast melee troop that targets buildings and can jump over the river. He followed the echoing call of "Hog Riderrrrr" all the way through the arena doors.`,
		attr.Elixir:  4,
		attr.BaseHP:  800,
		attr.DPS:     []interface{}{106, 117, 128, 141, 155, 170, 187, 205, 226, 248},
		attr.BaseDam: 160,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Buildings,
		attr.Speed:   speed.VeryFast,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:     "Wizard",
		attr.Arena:    arena.Arena5,
		attr.Rarity:   rarity.Rare,
		attr.Type:     typ.Troop,
		attr.Desc:     `The most awesome man to ever set foot in the arena, the Wizard will blow you away with his handsomeness... and/or fireballs.`,
		attr.Elixir:   5,
		attr.BaseHP:   340,
		attr.DPS:      []interface{}{76, 84, 92, 101, 111, 122, 134, 147, 161, 177},
		attr.BaseADam: 130,
		attr.HSpeed:   1.7,
		attr.Targets:  targets.AirAndGround,
		attr.Speed:    speed.Medium,
		attr.Range:    5.5,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:    "Three Musketeers",
		attr.Arena:   arena.Arena7,
		attr.Rarity:  rarity.Rare,
		attr.Type:    typ.Troop,
		attr.Desc:    `Trio of powerful, independent markswomen, fighting for justice and honor. Disrespecting them would not be just a mistake, it would be a cardinal sin!`,
		attr.Elixir:  9,
		attr.BaseHP:  340,
		attr.DPS:     []interface{}{90, 100, 110, 120, 132, 145, 160, 175, 192, 211},
		attr.BaseDam: 100,
		attr.HSpeed:  1.1,
		attr.Targets: targets.AirAndGround,
		attr.Speed:   speed.Medium,
		attr.Range:   6.5,
		attr.DTime:   1,
		attr.Count:   3,
	}.init(),

	// --- Epic Troops ---
	card{
		attr.Name:      "Witch",
		attr.Arena:     arena.Arena0,
		attr.Rarity:    rarity.Epic,
		attr.Type:      typ.Troop,
		attr.Desc:      `Summons skeletons, shoots destructo beams, has glowing pink eyes that unfortunately don't shoot lasers.`,
		attr.Elixir:    5,
		attr.BaseHP:    500,
		attr.DPS:       []interface{}{54, 58, 64, 70, 78, 85, 94, 104},
		attr.BaseADam:  38,
		attr.BaseSkeLV: 6,
		attr.SSpeed:    7.5,
		attr.HSpeed:    0.7,
		attr.Targets:   targets.AirAndGround,
		attr.Speed:     speed.Medium,
		attr.Range:     5.5,
		attr.DTime:     1,
	}.init(),
	card{
		attr.Name:    "Skeleton Army",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Epic,
		attr.Type:    typ.Troop,
		attr.Desc:    `Summons an army of skeletons. Meet Larry and his friends Harry, Gerry, Terry, Mary, etc.`,
		attr.Elixir:  4,
		attr.BaseHP:  30,
		attr.DPS:     []interface{}{30, 33, 36, 39, 43, 48, 52, 57},
		attr.BaseDam: 30,
		attr.HSpeed:  1,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Fast,
		attr.Range:   rng.Melee,
		attr.DTime:   1,
		attr.Count:   20,
	}.init(),
	card{
		attr.Name:     "Baby Dragon",
		attr.Arena:    arena.Arena0,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Troop,
		attr.Desc:     `Flying troop that deals area damage. Baby dragons hatch cute, hungry, and ready for a barbeque.`,
		attr.Elixir:   4,
		attr.BaseHP:   800,
		attr.DPS:      []interface{}{55, 61, 67, 73, 81, 88, 97, 107},
		attr.BaseADam: 100,
		attr.HSpeed:   1.8,
		attr.Targets:  targets.AirAndGround,
		attr.Speed:    speed.Fast,
		attr.Range:    3.5,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:    "Prince",
		attr.Arena:   arena.Arena0,
		attr.Rarity:  rarity.Epic,
		attr.Type:    typ.Troop,
		attr.Desc:    `Don't let the little pony fool you. Once the Prince gets a running start, you WILL be trampled. Does 2x damage once he gets charging.`,
		attr.Elixir:  5,
		attr.BaseHP:  1100,
		attr.DPS:     []interface{}{146, 161, 177, 194, 214, 234, 258, 282},
		attr.BaseDam: 220,
		attr.HSpeed:  1.5,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Medium,
		attr.Range:   2.5,
		attr.DTime:   1,
	}.init(),
	card{
		attr.Name:     "Giant Skeleton",
		attr.Arena:    arena.Arena2,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Troop,
		attr.Desc:     `The bigger the skeleton, the bigger the bomb. Carries a bomb that blows up when the Giant Skeleton dies.`,
		attr.Elixir:   6,
		attr.BaseHP:   2000,
		attr.DPS:      []interface{}{66, 73, 80, 88, 97, 106, 117, 128},
		attr.BaseDam:  100,
		attr.BaseDDam: 720,
		attr.HSpeed:   1.5,
		attr.Targets:  targets.Ground,
		attr.Speed:    speed.Medium,
		attr.Range:    rng.Melee,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:     "Balloon",
		attr.Arena:    arena.Arena2,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Troop,
		attr.Desc:     `As pretty as they are, you won't want a parade of THESE balloons showing up on the horizon. Drops powerful bombs and when shot down, crashes dealing area damage.`,
		attr.Elixir:   5,
		attr.BaseHP:   1050,
		attr.DPS:      []interface{}{200, 220, 242, 266, 292, 320, 352, 386},
		attr.BaseDam:  600,
		attr.BaseDDam: 100,
		attr.HSpeed:   3,
		attr.Targets:  targets.Buildings,
		attr.Speed:    speed.Medium,
		attr.Range:    rng.Melee,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:    "P.E.K.K.A",
		attr.Arena:   arena.Arena4,
		attr.Rarity:  rarity.Epic,
		attr.Type:    typ.Troop,
		attr.Desc:    `A heavily armored, slow melee fighter. Swings from the hip but packs a huge punch.`,
		attr.Elixir:  7,
		attr.BaseHP:  2600,
		attr.DPS:     []interface{}{250, 275, 302, 332, 365, 400, 440, 482},
		attr.BaseDam: 450,
		attr.HSpeed:  1.8,
		attr.Targets: targets.Ground,
		attr.Speed:   speed.Slow,
		attr.Range:   rng.Melee,
		attr.DTime:   3,
	}.init(),
	card{
		attr.Name:     "Golem",
		attr.Arena:    arena.Arena6,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Troop,
		attr.Desc:     `Slow but durable, only attacks buildings. When destroyed, explosively splits into two Golemites and deals area damage!`,
		attr.Elixir:   8,
		attr.BaseHP:   3000,
		attr.DPS:      []interface{}{74, 81, 90, 98, 108, 118, 130, 143},
		attr.BaseDam:  186,
		attr.BaseDDam: 186,
		attr.HSpeed:   2.5,
		attr.Targets:  targets.Buildings,
		attr.Speed:    speed.Slow,
		attr.Range:    rng.Melee,
		attr.DTime:    3,
	}.init(),
	card{
		attr.Name:     "Dark Prince",
		attr.Arena:    arena.Arena7,
		attr.Rarity:   rarity.Epic,
		attr.Type:     typ.Troop,
		attr.Desc:     `Dealing area damage with each swing and double after charging, the Dark Prince is a formidable fighter. To harm his squishy core, break his shield first.`,
		attr.Elixir:   4,
		attr.BaseHP:   700,
		attr.BaseSHP:  200,
		attr.DPS:      []interface{}{83, 91, 100, 110, 121, 133, 146, 159},
		attr.BaseADam: 125,
		attr.HSpeed:   1.5,
		attr.Targets:  targets.Ground,
		attr.Speed:    speed.Medium,
		attr.Range:    rng.Melee,
		attr.DTime:    1,
	}.init(),

	// --- Legendary Troops ---
	card{
		attr.Name:     "Ice Wizard",
		attr.Arena:    arena.Arena5,
		attr.Rarity:   rarity.Legendary,
		attr.Type:     typ.Troop,
		attr.Desc:     `This chill caster hails from the far North. He shoots ice shards at enemies, slowing down their movement and attack speed.`,
		attr.Elixir:   3,
		attr.BaseHP:   700,
		attr.DPS:      []interface{}{42, 46, 50, 55, 60, 67},
		attr.BaseADam: 63,
		attr.HSpeed:   1.5,
		attr.Targets:  targets.AirAndGround,
		attr.Speed:    speed.Medium,
		attr.Range:    6,
		attr.DTime:    1,
	}.init(),
	card{
		attr.Name:     "Princess",
		attr.Arena:    arena.Arena7,
		attr.Rarity:   rarity.Legendary,
		attr.Type:     typ.Troop,
		attr.Desc:     `Shoots a volley of flaming arrows halfway across the Arena. The Princess is afraid of germs, so keep the rabble away from her!`,
		attr.Elixir:   3,
		attr.BaseHP:   216,
		attr.DPS:      []interface{}{46, 51, 56, 62, 68, 74},
		attr.BaseADam: 140,
		attr.HSpeed:   3,
		attr.Targets:  targets.AirAndGround,
		attr.Speed:    speed.Medium,
		attr.Range:    9.5,
		attr.DTime:    1,
	}.init(),
}