#encoding "utf8"
#GRAMMAR_ROOT ROOT

Rent -> Word<kwset=[rent, populate]>;

Flat -> Word<kwset=[flat]> interp (+FactRent.Type="квартира");

AnyWordFlat -> AnyWord<kwset=~[rent, populate, studio, flat, room, neighbor, search, number, numeric]>;
AnyWordFlatNum -> AnyWordFlat<kwset=~[time, month]> {count=5};

A1 -> AnyWord<wff="(одно|ОДНО|евроодн|ЕВРООДН).*"> interp (+FactRent.Type="1");
A2 -> AnyWord<wff="(дву|ДВУ|евродву|ЕВРОДВУ).*"> interp (+FactRent.Type="2");
A3 -> AnyWord<wff="(тр.х|ТР.Х|евротр|ЕВРОТР).*"> interp (+FactRent.Type="3");
A4 -> AnyWord<wff="(четыр|ЧЕТЫР|евроче|ЕВРОЧЕ).*"> interp (+FactRent.Type="4");

RoomedTmp -> AnyWord<wff="(к|К).*|.*(ком|КОМ).*|.*(ю|Ю)|(ка|КА)|(ку|КУ)">;
Roomed -> RoomedTmp<kwset=~[flat, room, neighbor]>;

B1_1 -> AnyWord<wff="1\\D*"> Roomed interp (+FactRent.Type="1");
B1_2 -> AnyWord<wff="1\\D*"> AnyWordFlatNum Roomed interp (+FactRent.Type="1");
B1 -> B1_1 | B1_2;
B2_1 -> AnyWord<wff="2\\D*"> Roomed interp (+FactRent.Type="2");
B2_2 -> AnyWord<wff="2\\D*"> AnyWordFlatNum Roomed interp (+FactRent.Type="2");
B2 -> B2_1 | B2_2;
B3_1 -> AnyWord<wff="3\\D*"> Roomed interp (+FactRent.Type="3");
B3_2 -> AnyWord<wff="3\\D*"> AnyWordFlatNum Roomed interp (+FactRent.Type="3");
B3 -> B3_1 | B3_2;
B4_1 -> AnyWord<wff="[4-9]\\D*"> Roomed interp (+FactRent.Type="4");
B4_2 -> AnyWord<wff="[4-9]\\D*"> AnyWordFlatNum Roomed interp (+FactRent.Type="4");
B4 -> B4_1 | B4_2;

FlatNum -> A1 | A2 | A3 | A4 | B1 | B2 | B3 | B4;

C1 -> AnyWord<wff=".*однуш.*"> interp (+FactRent.Type="1 квартира");
C2 -> AnyWord<wff=".*двуш.*"> interp (+FactRent.Type="2 квартира");
C3 -> AnyWord<wff=".*тр.ш.*"> interp (+FactRent.Type="3 квартира");

FlatComplexTmp -> C1 | C2 | C3;
FlatComplex -> FlatComplexTmp<kwset=~[rent, populate, studio, flat, room, neighbor, search, number, numeric]>;

//сдам какую то квартиру
ROOT -> Rent AnyWordFlat* Flat { weight=0 };

//сдам какую то однушку
ROOT -> Rent AnyWordFlat* FlatComplex { weight=0.1 };

//сдам какую то 1 комнатную квартиру
ROOT -> Rent AnyWordFlat* FlatNum AnyWordFlat* Flat { weight=0.1 };

//сдам какую то квартиру 1 комнатную
ROOT -> Rent AnyWordFlat* Flat FlatNum { weight=0.1 };

Num -> AnyWord<wff="[1-5]"> interp (+FactRent.Type);
Flat3 -> AnyWordFlat Flat | Flat;
//сдам какаую то 1 квартиру
ROOT -> Rent AnyWordFlat* Num Flat3 { weight=0.1 };

//сдам какую то комнату
AnyWordRoom -> AnyWord<kwset=~[rent, populate, studio, flat, room, neighbor, search, flat_num]>;
Room -> Word<kwtype=[room]> interp (FactRent.Type="комната");
ROOT -> Rent AnyWordRoom* Room { weight=0.5 };

//сдам какую то студию
AnyWordStudio -> AnyWord<kwset=~[rent, populate, studio, neighbor, search, flat_num]>;
Studio -> Word<kwset=[studio]> interp (FactRent.Type="студия");
ROOT -> Rent AnyWordStudio* Studio { weight=1 };
