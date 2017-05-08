#encoding "utf8"
#GRAMMAR_ROOT ROOT

Rent -> Word<kwset=[rent_wrong, search, like]>;

AnyWordFlat -> AnyWord<kwset=~[rent, populate, studio, flat, room, neighbor, search, resident]>;
AnyWordFlatNum -> AnyWordFlat<kwset=~[time, month]> {count=5};

A1 -> AnyWord<wff="одно.*">;
A2 -> AnyWord<wff="дву.*">;
A3 -> AnyWord<wff="тр.x.*">;
A4 -> AnyWord<wff="четыр.*">;

RoomedTmp -> AnyWord<wff="(к|К).*|.*(ком|КОМ).*|.*(ю|Ю)|(ка|КА)|(ку|КУ)">;
Roomed -> RoomedTmp<kwset=~[flat, room, neighbor]>;

B1 -> AnyWord<wff="[1-9]\\D*"> Roomed;
B2 -> AnyWord<wff="[1-9]\\D*"> AnyWordFlatNum Roomed;

FlatNum -> B1 | B2;

Flat -> Word<kwset=[flat, area]> interp (+FactWrong.Wrong="flat");

ROOT -> Rent AnyWordFlat* FlatNum AnyWordFlat* Flat;

//------------- flat2

C1 -> AnyWord<wff=".*однуш.*">;
C2 -> AnyWord<wff=".*двуш.*">;
C3 -> AnyWord<wff=".*тр.ш.*">;

FlatComplexTmp -> C1 | C2 | C3;
FlatComplex -> FlatComplexTmp<kwset=~[flat, room, neighbor]>;

ROOT -> Rent AnyWordFlat* FlatComplex interp (FactWrong.Wrong="Flat_complex");

//---------------- flat3

ROOT -> Rent AnyWordFlat* Flat FlatNum;

//-------------- flat4

ROOT -> Rent AnyWordFlat* Flat;

//------------ house

House -> AnyWord<wff="(дом|коттедж|участо|мансард|жиль(e|ё)).*"> interp (FactWrong.Wrong="house");

ROOT -> Rent AnyWordFlat* House;

//------------- room

Room -> Word<kwtype=[room]> interp (FactWrong.Wrong="room");
AnyWordRoom -> AnyWord<kwset=~[rent, populate, studio, flat, room, neighbor, search, resident, flat_num]>;

ROOT -> Rent AnyWordRoom* Room;

//-------------- studio

Studio -> Word<kwset=[studio]> interp (FactWrong.Wrong="studio");
AnyWordStudio -> AnyWord<kwset=~[rent, populate, studio, neighbor, search, flat_num]>;

ROOT -> Rent AnyWordStudio* Studio;


