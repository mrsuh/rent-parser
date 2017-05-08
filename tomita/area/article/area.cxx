#encoding "utf8"
#GRAMMAR_ROOT ROOT

Delimeter -> Punct | Comma;
N1 -> AnyWord<wff="\\d{1,3}">;
N2 -> AnyWord<wff="\\d{1,3}"> Delimeter AnyWord<wff="\\d{1,3}">;
N3 -> AnyWord<wff="\\d{1,3}[\\.,]\\d{1,2}">;
Number -> N1 | N2 | N3;

//кв.м
M1_1 -> Word<wff="кв"> Punct Word<wff="м">;
M1_2 -> Word<wff="кв"> Word<wff="м">;
M1 -> M1_1 | M1_2;

//м2
M2 -> Word<wff="м"> AnyWord<wff="2">;

//м
M3 -> Word<wff="м">;

//метров
M4 -> Word<wff="метров">;

//м.кв
M5_1 -> Word<wff="м"> Punct Word<wff="кв">;
M5_2 -> Word<wff="м"> Word<wff="кв">;
M5 -> M5_1 | M5_2;

//м²
M6 -> AnyWord<wff="м"> '²';

//кв метра
M7_1 -> Word<wff="кв"> Punct Word<wff="метр.*">;
M7_2 -> Word<wff="кв"> Word<wff="метр.*">;
M7 -> M7_1 | M7_2;

Meter -> M1 | M2 | M3 | M4 | M5 | M6 | M7;

ROOT -> Number interp (FactArea.Area) Meter;