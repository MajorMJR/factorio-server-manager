package factorio

type DifficultySettings struct {
	RecipeDifficulty          int    `json:"recipe_difficulty"`
	TechnologyDifficulty      int    `json:"technology_difficulty"`
	TechnologyPriceMultiplier int    `json:"technology_price_multiplier"`
	ResearchQueueSetting      string `json:"research_queue_setting"`
}

type Pollution struct {
	Enabled                                 bool    `json:"enabled"`
	DiffusionRatio                          float32 `json:"diffusion_ratio"`
	MinToDiffuse                            float32 `json:"min_to_diffuse"`
	Ageing                                  float32 `json:"ageing"`
	ExpectedMaxPerChunk                     float32 `json:"expected_max_per_chunk"`
	MinToShowPerChunk                       float32 `json:"min_to_show_per_chunk"`
	MinPollutionToDamageTrees               float32 `json:"min_pollution_to_damage_trees"`
	PollutionWithMaxForestDamage            float32 `json:"pollution_with_max_forest_damage"`
	PollutionPerTreeDamage                  float32 `json:"pollution_per_tree_damage"`
	PollutionRestoredPerTreeDamage          float32 `json:"pollution_restored_per_tree_damage"`
	MaxPollutionToRestoreTrees              float32 `json:"max_pollution_to_restore_trees"`
	EnemyAttackPollutionConsumptionModifier float32 `json:"enemy_attack_pollution_consumption_modifier"`
}

type EnemyEvolution struct {
	Enabled         bool    `json:"enabled"`
	TimeFactor      float32 `json:"time_factor"`
	DestroyFactor   float32 `json:"destroy_factor"`
	PollutionFactor float32 `json:"pollution_factor"`
}

type EnemyExpansion struct {
	Enabled                          bool    `json:"enabled"`
	MinBaseSpacing                   int     `json:"min_base_spacing"`
	MaxExpansionDistance             int     `json:"max_expansion_distance"`
	FriendlyBaseInfluenceRadius      int     `json:"friendly_base_influence_radius"`
	EnemyBuildingInfluenceRadius     int     `json:"enemy_building_influence_radius"`
	BuildingCoefficient              float32 `json:"building_coefficient"`
	OtherBaseCoefficient             float32 `json:"other_base_coefficient"`
	NeighbouringChunkCoefficient     float32 `json:"neighbouring_chunk_coefficient"`
	NeighbouringBaseChunkCoefficient float32 `json:"neighbouring_base_chunk_coefficient"`
	MaxCollidingTilesCoefficient     float32 `json:"max_colliding_tiles_coefficient"`
	SettlerGroupMinSize              int     `json:"settler_group_min_size"`
	SettlerGroupMaxSize              int     `json:"settler_group_max_size"`
	MinExpansionCooldown             int     `json:"min_expansion_cooldown"`
	MaxExpansionCooldown             int     `json:"max_expansion_cooldown"`
}

type UnitGroup struct {
	MinGroupGatheringTime          int     `json:"min_group_gathering_time"`
	MaxGroupGatheringTime          int     `json:"max_group_gathering_time"`
	MaxWaitTimeForLateMembers      int     `json:"max_wait_time_for_late_members"`
	MaxGroupRadius                 float32 `json:"max_group_radius"`
	MinGroupRadius                 float32 `json:"min_group_radius"`
	MaxMemberSpeedupWhenBehind     float32 `json:"max_member_speedup_when_behind"`
	MaxMemberSlowdownWhenAhead     float32 `json:"max_member_slowdown_when_ahead"`
	MaxGroupSlowdownFactor         float32 `json:"max_group_slowdown_factor"`
	MaxGroupMemberFallbackFactor   int     `json:"max_group_member_fallback_factor"`
	MemberDisownDistance           int     `json:"member_disown_distance"`
	TickToleranceWhenMemberArrives int     `json:"tick_tolerance_when_member_arrives"`
	MaxGatheringUnitGroups         int     `json:"max_gathering_unit_groups"`
	MaxUnitGroupSize               int     `json:"max_unit_group_size"`
}

type SteeringConfig struct {
	Radius                     float32 `json:"radius"`
	SeparationForce            float32 `json:"separation_force"`
	SeparationFactor           float32 `json:"separation_factor"`
	ForceUnitFuzzyGotoBehavior bool    `json:"force_unit_fuzzy_goto_behavior"`
}

type Steering struct {
	Default SteeringConfig `json:"default"`
	Moving  SteeringConfig `json:"moving"`
}

type PathFinder struct {
	Fwd2bwdRatio                                  int     `json:"fwd2bwd_ratio"`
	GoalPressureRatio                             int     `json:"goal_pressure_ratio"`
	MaxStepsWorkedPerTick                         int     `json:"max_steps_worked_per_tick"`
	MaxWorkDonePerTick                            int     `json:"max_work_done_per_tick"`
	UsePathCache                                  bool    `json:"use_path_cache"`
	ShortCacheSize                                int     `json:"short_cache_size"`
	LongCacheSize                                 int     `json:"long_cache_size"`
	ShortCacheMinCacheableDistance                int     `json:"short_cache_min_cacheable_distance"`
	ShortCacheMinAlgoStepsToCache                 int     `json:"short_cache_min_algo_steps_to_cache"`
	LongCacheMinCacheableDistance                 int     `json:"long_cache_min_cacheable_distance"`
	CacheMaxConnectToCacheStepsMultiplier         int     `json:"cache_max_connect_to_cache_steps_multiplier"`
	CacheAcceptPathStartDistanceRatio             float32 `json:"cache_accept_path_start_distance_ratio"`
	CacheAcceptPathEndDistanceRatio               float32 `json:"cache_accept_path_end_distance_ratio"`
	NegativeCacheAcceptPathStartDistanceRatio     float32 `json:"negative_cache_accept_path_start_distance_ratio"`
	NegativeCacheAcceptPathEndDistanceRatio       float32 `json:"negative_cache_accept_path_end_distance_ratio"`
	CachePathStartDistanceRatingMultiplier        int     `json:"cache_path_start_distance_rating_multiplier"`
	CachePathEndDistanceRatingMultiplier          int     `json:"cache_path_end_distance_rating_multiplier"`
	StaleEnemyWithSameDestinationCollisionPenalty int     `json:"stale_enemy_with_same_destination_collision_penalty"`
	IgnoreMovingEnemyCollisionDistance            int     `json:"ignore_moving_enemy_collision_distance"`
	EnemyWithDifferentDestinationCollisionPenalty int     `json:"enemy_with_different_destination_collision_penalty"`
	GeneralEntityCollisionPenalty                 int     `json:"general_entity_collision_penalty"`
	GeneralEntitySubsequentCollisionPenalty       int     `json:"general_entity_subsequent_collision_penalty"`
	ExtendedCollisionPenalty                      int     `json:"extended_collision_penalty"`
	MaxClientsToAcceptAnyNewRequest               int     `json:"max_clients_to_accept_any_new_request"`
	MaxClientsToAcceptShortNewRequest             int     `json:"max_clients_to_accept_short_new_request"`
	DirectDistanceToConsiderShortRequest          int     `json:"direct_distance_to_consider_short_request"`
	ShortRequestMaxSteps                          int     `json:"short_request_max_steps"`
	ShortRequestRatio                             float32 `json:"short_request_ratio"`
	MinStepsToCheckPathFindTermination            int     `json:"min_steps_to_check_path_find_termination"`
	StartToGoalCostMultiplierToTerminatePathFind  float32 `json:"start_to_goal_cost_multiplier_to_terminate_path_find"`
	OverloadLevels                                []int   `json:"overload_levels"`
	OverloadMultipliers                           []int   `json:"overload_multipliers"`
	NegativePathCacheDelayInterval                int     `json:"negative_path_cache_delay_interval"`
}

type MapSettings struct {
	DifficultySettings     DifficultySettings `json:"difficulty_settings"`
	Pollution              Pollution          `json:"pollution"`
	EnemyEvolution         EnemyEvolution     `json:"enemy_evolution"`
	EnemyExpansion         EnemyExpansion     `json:"enemy_expansion"`
	UnitGroup              UnitGroup          `json:"unit_group"`
	Steering               Steering           `json:"steering"`
	PathFinder             PathFinder         `json:"path_finder"`
	MaxFailedBehaviorCount int                `json:"max_failed_behavior_count"`
}

func DefaultMapSettings() MapSettings {
	return MapSettings{
		DifficultySettings: DifficultySettings{
			RecipeDifficulty:          0,
			TechnologyDifficulty:      0,
			TechnologyPriceMultiplier: 1,
			ResearchQueueSetting:      "after-victory",
		},
		Pollution: Pollution{
			Enabled:                                 true,
			DiffusionRatio:                          0.02,
			MinToDiffuse:                            15,
			Ageing:                                  1,
			ExpectedMaxPerChunk:                     150,
			MinToShowPerChunk:                       50,
			MinPollutionToDamageTrees:               60,
			PollutionWithMaxForestDamage:            150,
			PollutionPerTreeDamage:                  50,
			PollutionRestoredPerTreeDamage:          10,
			MaxPollutionToRestoreTrees:              20,
			EnemyAttackPollutionConsumptionModifier: 1,
		},
		EnemyEvolution: EnemyEvolution{
			Enabled:         true,
			TimeFactor:      0.000004,
			DestroyFactor:   0.002,
			PollutionFactor: 0.0000009,
		},
		EnemyExpansion: EnemyExpansion{
			Enabled:                          true,
			MinBaseSpacing:                   3,
			MaxExpansionDistance:             7,
			FriendlyBaseInfluenceRadius:      2,
			EnemyBuildingInfluenceRadius:     2,
			BuildingCoefficient:              0.1,
			OtherBaseCoefficient:             2.0,
			NeighbouringChunkCoefficient:     0.5,
			NeighbouringBaseChunkCoefficient: 0.4,
			MaxCollidingTilesCoefficient:     0.9,
			SettlerGroupMinSize:              5,
			SettlerGroupMaxSize:              20,
			MinExpansionCooldown:             14400,
			MaxExpansionCooldown:             216000,
		},
		UnitGroup: UnitGroup{
			MinGroupGatheringTime:          3600,
			MaxGroupGatheringTime:          36000,
			MaxWaitTimeForLateMembers:      7200,
			MaxGroupRadius:                 30.0,
			MinGroupRadius:                 5.0,
			MaxMemberSpeedupWhenBehind:     1.4,
			MaxMemberSlowdownWhenAhead:     0.6,
			MaxGroupSlowdownFactor:         0.3,
			MaxGroupMemberFallbackFactor:   3,
			MemberDisownDistance:           10,
			TickToleranceWhenMemberArrives: 60,
			MaxGatheringUnitGroups:         30,
			MaxUnitGroupSize:               200,
		},
		Steering: Steering{
			Default: SteeringConfig{
				Radius:                     1.2,
				SeparationForce:            0.005,
				SeparationFactor:           1.2,
				ForceUnitFuzzyGotoBehavior: false,
			},
			Moving: SteeringConfig{
				Radius:                     3,
				SeparationForce:            0.01,
				SeparationFactor:           3,
				ForceUnitFuzzyGotoBehavior: false,
			},
		},
		PathFinder: PathFinder{
			Fwd2bwdRatio:                                  5,
			GoalPressureRatio:                             2,
			MaxStepsWorkedPerTick:                         100,
			MaxWorkDonePerTick:                            8000,
			UsePathCache:                                  true,
			ShortCacheSize:                                5,
			LongCacheSize:                                 25,
			ShortCacheMinCacheableDistance:                10,
			ShortCacheMinAlgoStepsToCache:                 50,
			LongCacheMinCacheableDistance:                 30,
			CacheMaxConnectToCacheStepsMultiplier:         1000,
			CacheAcceptPathStartDistanceRatio:             0.2,
			CacheAcceptPathEndDistanceRatio:               0.15,
			NegativeCacheAcceptPathStartDistanceRatio:     0.3,
			NegativeCacheAcceptPathEndDistanceRatio:       0.3,
			CachePathStartDistanceRatingMultiplier:        10,
			CachePathEndDistanceRatingMultiplier:          20,
			StaleEnemyWithSameDestinationCollisionPenalty: 30,
			IgnoreMovingEnemyCollisionDistance:            5,
			EnemyWithDifferentDestinationCollisionPenalty: 30,
			GeneralEntityCollisionPenalty:                 10,
			GeneralEntitySubsequentCollisionPenalty:       3,
			ExtendedCollisionPenalty:                      3,
			MaxClientsToAcceptAnyNewRequest:               10,
			MaxClientsToAcceptShortNewRequest:             100,
			DirectDistanceToConsiderShortRequest:          100,
			ShortRequestMaxSteps:                          1000,
			ShortRequestRatio:                             0.5,
			MinStepsToCheckPathFindTermination:            2000,
			StartToGoalCostMultiplierToTerminatePathFind:  500.0,
			OverloadLevels:                                []int{0, 100, 500},
			OverloadMultipliers:                           []int{2, 3, 4},
			NegativePathCacheDelayInterval:                20,
		},
		MaxFailedBehaviorCount: 3,
	}
}
