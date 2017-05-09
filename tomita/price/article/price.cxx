#encoding "utf8"
#GRAMMAR_ROOT ROOT

NumberFull -> Word<wff="\\d{5,6}"> interp (FactPrice.Full);
NumberShort -> Word<wff="\\d{1,2}"> interp (FactPrice.Short);
NumberHalf -> Word<wff="\\d{1,2}([^а-яА-Я\\d])\\d{1,3}"> interp (FactPrice.Half);
Number -> NumberFull | NumberShort | NumberHalf;

P1 -> AnyWord<wff="вместе"> AnyWord<wff="с">;
P2 -> AnyWord<wff="(\\+|и)">;
Plus -> P1 | P2;

Price -> AnyWord<wff="цена|оплата">;

AllInclude -> AnyWord<wff="всё"> AnyWord<wff="включе(н|нн)о">;

C1 -> Word<wff="ку">;
C2 -> Word<wff="к"> AnyWord<wff=/(\.|\/)/> Word<wff="у">;
C3 -> Word<wff="свет">;

Communal -> C1 | C2 | C3;

PL1 -> Word<wff="без"> Word<wff="залога">;
PL2 -> Plus Word<wff="залог">;
Deposit -> PL1 | PL2;

InMonth -> AnyWord<wff=/(\/|в)/> AnyWord<wff="мес.*">;

CU1 -> AnyWord<wff="(руб|рублей)">;
CU2 -> AnyWord<wff="(т|тысяч|тыс)"> AnyWord<wff=/(\.|\/)/> AnyWord<wff="(р|рублей|руб)">;
CU3 -> AnyWord<wff="(т|тысяч|тыс)"> AnyWord<wff="(р|рублей|руб)">;
CU4 -> AnyWord<wff="(р|рублей|руб)"> AnyWord<wff="\\.">;
CU5 -> AnyWord<wff="(р|рублей|руб)">;
CU6 -> AnyWord<wff="(т|тысяч|тыс)">;
CU7 -> AnyWord<wff="(тр)">;

Currency -> CU1 | CU2 | CU3 | CU4 | CU5 | CU6 | CU7;

ROOT -> Price Number {weight=0.5};
ROOT -> Number Plus Communal {weight=1};
ROOT -> Number AllInclude {weight=0.5};
ROOT -> Number InMonth {weight=0.5};
ROOT -> Number Deposit {weight=1};

ROOT -> Price Number Currency {weight=0.5};
ROOT -> Number Currency {weight=0.1};
ROOT -> Number Currency Plus Communal {weight=1};
ROOT -> Number Currency AllInclude {weight=0.5};
ROOT -> Number Currency InMonth {weight=0.5};
ROOT -> Number Currency Deposit {weight=1};

