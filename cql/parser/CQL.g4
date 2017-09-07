
/** Derived from github.com/RedisLabsModules/secondary/docs/Commands.md

https://stackoverflow.com/questions/44796556/how-to-tell-if-antlr4-parser-consumed-all-tokens-of-the-tokenstream
 */

grammar CQL;

cql
    : create EOF
    | destroy EOF
    | insert EOF
    | del EOF
    | query EOF
    ;

create: 'IDX.CREATE' indexName 'SCHEMA' (uintPropDef)* (enumPropDef)* (strPropDef)*;

destroy: 'IDX.DESTROY' indexName;

insert: 'IDX.INSERT' document;

del: 'IDX.DEL' document;

query: ('IDX.SELECT' | 'QUERY') indexName 'WHERE' (uintPred)* (enumPred)* (strPred)* orderLimit?;

indexName: IDENTIFIER;

document: indexName docId value+;

uintPropDef: property uintType;

enumPropDef: property K_ENUM;

strPropDef: property K_STRING;

orderLimit: 'ORDERBY' order ('LIMIT' limit)?;

order: property;

property: IDENTIFIER;

uintType
    : K_UINT8
    | K_UINT16
    | K_UINT32
    | K_UINT64
    ;

docId: INT;

value
    : INT
    | STRING
    ;

uintPred: property compare INT;

enumPred: property K_IN intList;

strPred: property K_CONTAINS STRING;

compare
    : K_LT
    | K_BT
    | K_EQ
    | K_LE
    | K_BE
    ;

intList: '[' INT (',' INT)* ']';

limit: INT;

K_UINT8: 'UINT8';
K_UINT16: 'UINT16';
K_UINT32: 'UINT32';
K_UINT64: 'UINT64';
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


INT
    : '0' | [1-9] [0-9]*
    ;


// no leading zeros

fragment EXP
    : [Ee] [+\-]? INT
    ;


IDENTIFIER
    : [a-zA-Z_]([a-zA-Z0-9_])*
    ;

// \- since - means "range" inside [...]

WS
    : [ \t\n\r] + -> skip
    ;

