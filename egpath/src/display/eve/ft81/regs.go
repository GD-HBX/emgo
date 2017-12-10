package ft81

// Register addresses.
const (
	REG_ID         = RAM_REG + 0
	REG_FRAMES     = RAM_REG + 4
	REG_CLOCK      = RAM_REG + 8
	REG_FREQUENCY  = RAM_REG + 12
	REG_RENDERMODE = RAM_REG + 16
	REG_SNAPY      = RAM_REG + 20
	REG_SNAPSHOT   = RAM_REG + 24
	REG_SNAPFORMAT = RAM_REG + 28
	REG_CPURESET   = RAM_REG + 32
	REG_TAP_CRC    = RAM_REG + 36
	REG_TAP_MASK   = RAM_REG + 40
	REG_HCYCLE     = RAM_REG + 44
	REG_HOFFSET    = RAM_REG + 48
	REG_HSIZE      = RAM_REG + 52
	REG_HSYNC0     = RAM_REG + 56
	REG_HSYNC1     = RAM_REG + 60
	REG_VCYCLE     = RAM_REG + 64
	REG_VOFFSET    = RAM_REG + 68
	REG_VSIZE      = RAM_REG + 72
	REG_VSYNC0     = RAM_REG + 76
	REG_VSYNC1     = RAM_REG + 80
	REG_DLSWAP     = RAM_REG + 84
	REG_ROTATE     = RAM_REG + 88
	REG_OUTBITS    = RAM_REG + 92
	REG_DITHER     = RAM_REG + 96
	REG_SWIZZLE    = RAM_REG + 100
	REG_CSPREAD    = RAM_REG + 104
	REG_PCLK_POL   = RAM_REG + 108
	REG_PCLK       = RAM_REG + 112
	REG_TAG_X      = RAM_REG + 116
	REG_TAG_Y      = RAM_REG + 120
	REG_TAG        = RAM_REG + 124
	REG_VOL_PB     = RAM_REG + 128
	REG_VOL_SOUND  = RAM_REG + 132
	REG_SOUND      = RAM_REG + 136
	REG_PLAY       = RAM_REG + 140
	REG_GPIO_DIR   = RAM_REG + 144
	REG_GPIO       = RAM_REG + 148
	REG_GPIOX_DIR  = RAM_REG + 152
	REG_GPIOX      = RAM_REG + 156

	REG_INT_FLAGS        = RAM_REG + 168
	REG_INT_EN           = RAM_REG + 172
	REG_INT_MASK         = RAM_REG + 176
	REG_PLAYBACK_START   = RAM_REG + 180
	REG_PLAYBACK_LENGTH  = RAM_REG + 184
	REG_PLAYBACK_READPTR = RAM_REG + 188
	REG_PLAYBACK_FREQ    = RAM_REG + 192
	REG_PLAYBACK_FORMAT  = RAM_REG + 196
	REG_PLAYBACK_LOOP    = RAM_REG + 200
	REG_PLAYBACK_PLAY    = RAM_REG + 204
	REG_PWM_HZ           = RAM_REG + 208
	REG_PWM_DUTY         = RAM_REG + 212
	REG_MACRO_0          = RAM_REG + 216
	REG_MACRO_1          = RAM_REG + 220

	REG_BUSYBITS = RAM_REG + 232

	REG_ROMSUB_SEL = RAM_REG + 240

	REG_CMD_READ            = RAM_REG + 248
	REG_CMD_WRITE           = RAM_REG + 252
	REG_CMD_DL              = RAM_REG + 256
	REG_TOUCH_MODE          = RAM_REG + 260
	REG_CTOUCH_EXTENDED     = RAM_REG + 264
	REG_TOUCH_ADC_MODE      = RAM_REG + 264
	REG_TOUCH_CHARGE        = RAM_REG + 268
	REG_TOUCH_SETTLE        = RAM_REG + 272
	REG_TOUCH_OVERSAMPLE    = RAM_REG + 276
	REG_TOUCH_RZTHRESH      = RAM_REG + 280
	REG_CTOUCH_TOUCH1_XY    = RAM_REG + 284
	REG_TOUCH_RAW_XY        = RAM_REG + 284
	REG_CTOUCH_TOUCH4_Y     = RAM_REG + 288
	REG_TOUCH_RZ            = RAM_REG + 288
	REG_CTOUCH_TOUCH0_XY    = RAM_REG + 292
	REG_TOUCH_SCREEN_XY     = RAM_REG + 292
	REG_TOUCH_TAG_XY        = RAM_REG + 296
	REG_TOUCH_TAG           = RAM_REG + 300
	REG_TOUCH_TAG1_XY       = RAM_REG + 304
	REG_TOUCH_TAG1          = RAM_REG + 308
	REG_TOUCH_TAG2_XY       = RAM_REG + 312
	REG_TOUCH_TAG2          = RAM_REG + 316
	REG_TOUCH_TAG3_XY       = RAM_REG + 320
	REG_TOUCH_TAG3          = RAM_REG + 324
	REG_TOUCH_TAG4_XY       = RAM_REG + 328
	REG_TOUCH_TAG4          = RAM_REG + 332
	REG_TOUCH_TRANSFORM_A   = RAM_REG + 336
	REG_TOUCH_TRANSFORM_B   = RAM_REG + 340
	REG_TOUCH_TRANSFORM_C   = RAM_REG + 344
	REG_TOUCH_TRANSFORM_D   = RAM_REG + 348
	REG_TOUCH_TRANSFORM_E   = RAM_REG + 352
	REG_TOUCH_TRANSFORM_F   = RAM_REG + 356
	REG_CYA_TOUCH           = RAM_REG + 360
	REG_ANALOG              = RAM_REG + 364
	REG_CTOUCH_TOUCH4_X     = RAM_REG + 364
	REG_PATCHED_TOUCH_FAULT = RAM_REG + 364
	REG_PATCHED_ANALOG      = RAM_REG + 368
	REG_TOUCH_FAULT         = RAM_REG + 368
	REG_BIST_EN             = RAM_REG + 372
	REG_CRC                 = RAM_REG + 376
	REG_SPI_EARLY_TX        = RAM_REG + 380
	REG_TRIM                = RAM_REG + 384
	REG_ANA_COMP            = RAM_REG + 388
	REG_SPI_WIDTH           = RAM_REG + 392
	REG_CTOUCH_TOUCH2_XY    = RAM_REG + 396
	REG_TOUCH_DIRECT_XY     = RAM_REG + 396
	REG_CTOUCH_TOUCH3_XY    = RAM_REG + 400
	REG_TOUCH_DIRECT_Z1Z2   = RAM_REG + 400

	REG_DATESTAMP       = RAM_REG + 1380
	REG_CMDB_SPACE      = RAM_REG + 1396
	REG_CMDB_WRITE      = RAM_REG + 1400
	
	REG_TRACKER         = RAM_REG + 28672
	REG_TRACKER_1       = RAM_REG + 28676
	REG_TRACKER_2       = RAM_REG + 28680
	REG_TRACKER_3       = RAM_REG + 28684
	REG_TRACKER_4       = RAM_REG + 28688
	REG_MEDIAFIFO_READ  = RAM_REG + 28692
	REG_MEDIAFIFO_WRITE = RAM_REG + 28696
)
