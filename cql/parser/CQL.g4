
/** Derived from github.com/RedisLabsModules/secondary/docs/Commands.md */

grammar CQL;

cql
    : create
    | destroy
    | insert
    | del
    | query
    ;

create: 'IDX.CREATE' indexName 'SCHEMA' (propertyDef)+;

destroy: 'IDX.DESTROY' indexName;

insert: 'IDX.INSERT' indexName docId value+;

del: 'IDX.DEL' indexName docId value+;

query: 'IDX.SELECT' indexName 'WHERE' predicates 'ORDERBY' property ('LIMIT' limit)?;

indexName: Identifier;

propertyDef: property popType;

property: Identifier;

popType:
    | K_UINT8
    | K_UINT16
    | K_UINT32
    | K_UINT64
    | K_FLOAT
    | K_ENUM
    | K_STRING
    ;

docId: INT;

value
    : NUMBER
    | STRING
    ;

predicates: predicate+;

predicate: property relate value;

relate
    : compare
    | K_IN
    | K_CONTAINS
    ;

compare
    : K_LT
    | K_BT
    | K_EQ
    | K_LE
    | K_BE
    ;

limit: INT;

K_UINT8: 'UINT8';
K_UINT16: 'UINT16';
K_UINT32: 'UINT32';
K_UINT64: 'UINT64';
K_FLOAT: 'FLOAT';
K_ENUM: 'ENUM';
K_STRING: 'STRING';
K_IN: 'IN';
K_CONTAINS: 'CONTAINS';
K_LT: '<';
K_BT: '>';
K_EQ: '=';
K_LE: '<=';
K_BE: '>=';

STRING
    : '"' (ESC | ~ ["\\])* '"'
    ;


fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;


fragment UNICODE
    : 'u' HEX HEX HEX HEX
    ;


fragment HEX
    : [0-9a-fA-F]
    ;


NUMBER
    : '-'? INT '.' [0-9] + EXP? | '-'? INT EXP | '-'? INT
    ;


INT
    : '0' | [1-9] [0-9]*
    ;

// no leading zeros

fragment EXP
    : [Ee] [+\-]? INT
    ;


Identifier
    :   IdentifierNondigit
        (   IdentifierNondigit
        |   Digit
        )*
    ;

fragment
IdentifierNondigit
    : Nondigit
    ;

fragment
Nondigit
    :   [a-zA-Z_]
    ;

fragment
Digit
    :   [0-9]
    ;

// \- since - means "range" inside [...]

WS
    : [ \t\n\r] + -> skip
    ;

