#encoding "utf8"
#GRAMMAR_ROOT ROOT

//89992144342
ROOT -> AnyWord<wff="[78]\\d{10}"> interp (FactContact.Contact);

//9992144342
ROOT -> AnyWord<wff="\\d{10}"> interp (FactContact.Contact);