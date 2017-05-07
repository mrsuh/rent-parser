#encoding "utf8"
#GRAMMAR_ROOT ROOT

//---------------- qua

Rent -> Word<kwset=[rent, populate]>;
Question -> AnyWord<wff="(\\?)"> interp (FactError.Error="question");

ROOT -> Rent Question;

//---------------- vk

Quote -> AnyWord<wff="\\[">;
Id -> AnyWord<wff="id"> interp (FactError.Error="vk");
No -> AnyWord<wff="\\d*:\\w*">;

ROOT -> Quote Id No;