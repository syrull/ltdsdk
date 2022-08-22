CREATE TABLE games (
	id INTEGER PRIMARY KEY,
	gameId TEXT NOT NULL,
	version TEXT NOT NULL,
	date DATE NOT NULL,
	queue_type TEXT NOT NULL,
	ending_wave INTEGER NOT NULL,
	game_length INTEGER NOT NULL,
	game_elo INTEGER NOT NULL,
	player_count INTEGER NOT NULL,
	human_count INTEGER NOT NULL,
	spell_choices_csv TEXT NOT NULL,
	left_king_percent_hp_csv TEXT NOT NULL,
	right_king_percent_hp_csv TEXT NOT NULL,
	king_spell TEXT NOT NULL
);