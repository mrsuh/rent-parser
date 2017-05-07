#encoding "utf8"
#GRAMMAR_ROOT ROOT

Search -> Word<kwset=[search]>;
Neighbor -> Word<kwset=[resident]>;

AnyWordFlat -> AnyWord<kwset=~[rent, populate, studio, flat, room, search, numeric]>;
AnyWordFlatNum -> AnyWordFlat<kwset=~[time, month]> {count=5};
AnyWordFlatSeveral -> AnyWordFlat {count=5};

A1 -> AnyWord<wff="(одно|ОДНО|евроодн|ЕВРООДН).*"> interp (+FactNeighbor.Type="1");
A2 -> AnyWord<wff="(дву|ДВУ|евродву|ЕВРОДВУ).*"> interp (+FactNeighbor.Type="2");
A3 -> AnyWord<wff="(тр.х|ТР.Х|евротр|ЕВРОТР).*"> interp (+FactNeighbor.Type="3");
A4 -> AnyWord<wff="(четыр|ЧЕТЫР|евроче|ЕВРОЧЕ).*"> interp (+FactNeighbor.Type="4");

RoomedTmp -> AnyWord<wff="(к|К).*|.*(ком|КОМ).*|.*(ю|Ю)|(ка|КА)|(ку|КУ)">;
Roomed -> RoomedTmp<kwset=~[flat, room, neighbor]>;

B1_1 -> AnyWord<wff="1\\D*"> Roomed interp (+FactNeighbor.Type="1");
B1_2 -> AnyWord<wff="1\\D*"> AnyWordFlatNum Roomed interp (+FactNeighbor.Type="1");
B1 -> B1_1 | B1_2;
B2_1 -> AnyWord<wff="2\\D*"> Roomed interp (+FactNeighbor.Type="2");
B2_2 -> AnyWord<wff="2\\D*"> AnyWordFlatNum Roomed interp (+FactNeighbor.Type="2");
B2 -> B2_1 | B2_2;
B3_1 -> AnyWord<wff="3\\D*"> Roomed interp (+FactNeighbor.Type="3");
B3_2 -> AnyWord<wff="3\\D*"> AnyWordFlatNum Roomed interp (+FactNeighbor.Type="3");
B3 -> B3_1 | B3_2;
B4_1 -> AnyWord<wff="[4-9]\\D*"> Roomed interp (+FactNeighbor.Type="4");
B4_2 -> AnyWord<wff="[4-9]\\D*"> AnyWordFlatNum Roomed interp (+FactNeighbor.Type="4");
B4 -> B4_1 | B4_2;

FlatNum -> A1 | A2 | A3 | A4 | B1 | B2 | B3 | B4;

Flat -> Word<kwset=[flat]> interp (+FactNeighbor.Type="квартира");

ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* FlatNum AnyWordFlat* Flat { weight=0.1 };

//---------------- flat2

C1 -> AnyWord<wff=".*однуш.*"> interp (+FactNeighbor.Type="1 квартира");
C2 -> AnyWord<wff=".*двуш.*"> interp (+FactNeighbor.Type="2 квартира");
C3 -> AnyWord<wff=".*тр.ш.*"> interp (+FactNeighbor.Type="3 квартира");

FlatComplexTmp -> C1 | C2 | C3;
FlatComplex -> FlatComplexTmp<kwset=~[flat, room, neighbor]>;

ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* FlatComplex { weight=0.1 };

//------------------ flat3

ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* Flat FlatNum { weight=0.1 };
ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* FlatNum Flat { weight=0.1 };

//--------------------- flat4

ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* Flat { weight=0 };

//-------------------- flat5

Num -> AnyWord<wff="[1-5]"> interp (+FactNeighbor.Type);

FF -> AnyWordFlatSeveral Flat | Flat;

ROOT -> Search AnyWordFlat* Neighbor AnyWordFlat* Num FF { weight=0.1 };

//------------------------ room

Room -> Word<kwtype=[room]> interp (FactNeighbor.Type="комната");
AnyWordRoom -> AnyWord<kwset=~[rent, populate, studio, flat, room, search, flat_num]>;

ROOT -> Search AnyWordRoom* Neighbor AnyWordRoom* Room { weight=0.5 };

//-------------------- studio

Studio -> Word<kwset=[studio]> interp (FactNeighbor.Type="студия");
AnyWordStudio -> AnyWord<kwset=~[rent, populate, studio, neighbor, search, flat_num]>;

ROOT -> Search AnyWordStudio* Neighbor AnyWordStudio* Studio { weight=1 };
