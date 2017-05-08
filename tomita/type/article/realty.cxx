#encoding "utf8"
#GRAMMAR_ROOT ROOT

AnyWordFlat -> AnyWord<kwset=~[rent, populate, studio, flat, room, neighbor, search, numeric]>;
AnyWordFlatNum -> AnyWordFlat<kwset=~[time, month]> {count=5};

A1 -> AnyWord<wff="(одно|ОДНО|евроодн|ЕВРООДН).*"> interp (+FactRealty.Type="1");
A2 -> AnyWord<wff="(дву|ДВУ|евродву|ЕВРОДВУ).*"> interp (+FactRealty.Type="2");
A3 -> AnyWord<wff="(тр.х|ТР.Х|евротр|ЕВРОТР).*"> interp (+FactRealty.Type="3");
A4 -> AnyWord<wff="(четыр|ЧЕТЫР|евроче|ЕВРОЧЕ).*"> interp (+FactRealty.Type="4");

RoomedTmp -> AnyWord<wff="(к|К).*|.*(ком|КОМ).*|.*(ю|Ю)|(ка|КА)|(ку|КУ)">;
Roomed -> RoomedTmp<kwset=~[flat, room, neighbor]>;

B1_1 -> AnyWord<wff="1\\D*"> Roomed interp (+FactRealty.Type="1");
B1_2 -> AnyWord<wff="1\\D*"> AnyWordFlatNum Roomed interp (+FactRealty.Type="1");
B1 -> B1_1 | B1_2;
B2_1 -> AnyWord<wff="2\\D*"> Roomed interp (+FactRealty.Type="2");
B2_2 -> AnyWord<wff="2\\D*"> AnyWordFlatNum Roomed interp (+FactRealty.Type="2");
B2 -> B2_1 | B2_2;
B3_1 -> AnyWord<wff="3\\D*"> Roomed interp (+FactRealty.Type="3");
B3_2 -> AnyWord<wff="3\\D*"> AnyWordFlatNum Roomed interp (+FactRealty.Type="3");
B3 -> B3_1 | B3_2;
B4_1 -> AnyWord<wff="4\\D*"> Roomed interp (+FactRealty.Type="4");
B4_2 -> AnyWord<wff="4\\D*"> AnyWordFlatNum Roomed interp (+FactRealty.Type="4");
B4 -> B4_1 | B4_2;

FlatNum -> A1 | A2 | A3 | A4 | B1 | B2 | B3 | B4;

Flat -> Word<kwset=[flat]> interp (+FactRealty.Type="квартира");

ROOT -> FlatNum AnyWordFlat* Flat { weight=0.1 };

//------------------------ flat2

C1 -> AnyWord<wff=".*однуш.*"> interp (+FactRealty.Type="1 квартира");
C2 -> AnyWord<wff=".*двуш.*"> interp (+FactRealty.Type="2 квартира");
C3 -> AnyWord<wff=".*тр.ш.*"> interp (+FactRealty.Type="3 квартира");

FlatComplexTmp -> C1 | C2 | C3;
FlatComplex -> FlatComplexTmp<kwset=~[flat, room, neighbor]>;

ROOT -> FlatComplex { weight=0.1 };

//------------------------ flat3

ROOT -> Flat FlatNum { weight=0.1 };

//---------------------- flat4

ROOT -> Flat { weight=0 };

//----------------------- room

ROOT -> Word<kwtype=[room]> interp (FactRealty.Type="комната") { weight=0.5 };

//------------------------ studio

ROOT -> Word<kwset=[studio]> interp (FactRealty.Type="студия") { weight=1 };
