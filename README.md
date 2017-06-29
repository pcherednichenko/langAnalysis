### About the project:

Creates reports on the analysis of the popularity of various programming languages for different cities in Russia

The data is taken from HH Api (the most popular site for job search)

Also analyzes the popularity of programming languages in GitHub

To start the project:
```
docker-compose up -d
```

After starting the project will be available on the:
```
localhost:8090/
```

**Project at the initial stage of development**

An example of what the program will output (for June 29, 2017):
```
app_1       | GitHub language analysis:
app_1       |   Erlang  :  19853
app_1       |   Dart  :  8471
app_1       |   Visual Basic  :  179213
app_1       |   CoffeeScript  :  64753
app_1       |   Lua  :  83623
app_1       |   Objective-C  :  411216
app_1       |   Swift  :  278484
app_1       |   Fortran  :  10612
app_1       |   Haskell  :  63238
app_1       |   C#  :  706748
app_1       |   Scala  :  103067
app_1       |   R  :  214474
app_1       |   PHP  :  1267551
app_1       |   Perl  :  129361
app_1       |   C++  :  706748
app_1       |   HTML  :  1929298
app_1       |   Matlab  :  87572
app_1       |   JavaScript  :  3773193
app_1       |   Shell  :  528707
app_1       |   Go  :  211339
app_1       |   TeX  :  89241
app_1       |   Python  :  1783178
app_1       |   Java  :  3331112
app_1       |   Ruby  :  1280860
 
app_1       | Москва (Moscow)
app_1       |    Dart  :  2
app_1       |    Python  :  1230
app_1       |    Visual Basic  :  101
app_1       |    Scala  :  124
app_1       |    Fortran  :  7
app_1       |    Go  :  73
app_1       |    Swift  :  161
app_1       |    Matlab  :  111
app_1       |    JavaScript  :  2221
app_1       |    Haskell  :  12
app_1       |    CoffeeScript  :  26
app_1       |    TeX  :  5
app_1       |    Ruby  :  273
app_1       |    Lua  :  42
app_1       |    C++  :  938
app_1       |    Objective-C  :  167
app_1       |    Shell  :  243
app_1       |    C#  :  948
app_1       |    PHP  :  1202
app_1       |    Erlang  :  18
app_1       |    Perl  :  240
app_1       |    HTML  :  2093
app_1       |    Java  :  1508
 
app_1       | Санкт-Петербург (Saint-Petersburg)
app_1       |    C++  :  500
app_1       |    Shell  :  104
app_1       |    C#  :  381
app_1       |    Perl  :  89
app_1       |    TeX  :  1
app_1       |    Haskell  :  11
app_1       |    Lua  :  18
app_1       |    PHP  :  502
app_1       |    Go  :  27
app_1       |    Matlab  :  55
app_1       |    Ruby  :  107
app_1       |    Scala  :  64
app_1       |    CoffeeScript  :  18
app_1       |    Fortran  :  3
app_1       |    Java  :  724
app_1       |    Python  :  505
app_1       |    Objective-C  :  74
app_1       |    Visual Basic  :  30
app_1       |    Erlang  :  28
app_1       |    JavaScript  :  955
app_1       |    Swift  :  57
app_1       |    HTML  :  859
app_1       |    Dart  :  11
```