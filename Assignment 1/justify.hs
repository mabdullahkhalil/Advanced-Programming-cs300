{-    README

Name: Muhammad Abdullah Khalil
Roll: 19100142

EVERYTHING WORKS!!! (well hopefully)


-}


















-- ========================================================================================================================== --


--
--                                                          ASSIGNMENT 1
--
--      A common type of text alignment in print media is "justification", where the spaces between words, are stretched or
--      compressed to align both the left and right ends of each line of text. In this problem we'll be implementing a text
--      justification function for a monospaced terminal output (i.e. fixed width font where every letter has the same width).
--
--      Alignment is achieved by inserting blanks and hyphenating the words. For example, given a text:
--
--              "He who controls the past controls the future. He who controls the present controls the past."
--
--      we want to be able to align it like this (to a width of say, 15 columns):
--
--              He who controls
--              the  past cont-
--              rols  the futu-
--              re. He  who co-
--              ntrols the pre-
--              sent   controls
--              the past.
--


-- ========================================================================================================================== --



import Data.List
import Data.Char
import Data.List.Split

text1 = "He who controls the past controls the future. He who controls the present controls the past."
text2 = "A creative man is motivated by the desire to achieve, not by the desire to beat others."


-- ========================================================================================================================== --







-- ========================================================= PART 1 ========================================================= --


--
-- Define a function that splits a list of words into two lists, such that the first list does not exceed a given line width.
-- The function should take an integer and a list of words as input, and return a pair of lists.
-- Make sure that spaces between words are counted in the line width.
--
-- Example:
--    splitLine ["A", "creative", "man"] 12   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 11   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 10   ==>   (["A", "creative"], ["man"])
--    splitLine ["A", "creative", "man"] 9    ==>   (["A"], ["creative", "man"])
--



splitl::[String]->Int->[String]
splitl _ 0=[]
splitl [] _=[]
splitl (x:xs) y | (y-(length x)>=0)=  x:(splitl xs (y-(length x)-1))
                | otherwise = []
                

splitll::[String]->[String]->[String]
splitll [] []=[]
splitll [] (y:ys) = y:(splitll [] ys)
splitll (x:xs) (y:ys) | x==y = (splitll xs ys)



splitLine :: [String] -> Int -> ([String], [String])
splitLine x y= (splitl x y,((splitll ((splitl x) y)) x))






-- ========================================================= PART 2 ========================================================= --


--
-- To be able to align the lines nicely. we have to be able to hyphenate long words. Although there are rules for hyphenation
-- for each language, we will take a simpler approach here and assume that there is a list of words and their proper hyphenation.
-- For example:

enHyp = [("creative", ["cr","ea","ti","ve"]), ("controls", ["co","nt","ro","ls"]), ("achieve", ["ach","ie","ve"]), ("future", ["fu","tu","re"]), ("present", ["pre","se","nt"]), ("motivated", ["mot","iv","at","ed"]), ("desire", ["de","si","re"]), ("others", ["ot","he","rs"])]


--
-- Define a function that splits a list of words into two lists in different ways. The first list should not exceed a given
-- line width, and may include a hyphenated part of a word at the end. You can use the splitLine function and then attempt
-- to breakup the next word with a given list of hyphenation rules. Include a breakup option in the output only if the line
-- width constraint is satisfied.
-- The function should take a hyphenation map, an integer line width and a list of words as input. Return pairs of lists as
-- in part 1.
--
-- Example:
--    lineBreaks enHyp 12 ["He", "who", "controls."]   ==>   [(["He","who"], ["controls."]), (["He","who","co-"], ["ntrols."]), (["He","who","cont-"], ["rols."])]
--
-- Make sure that words from the list are hyphenated even when they have a trailing punctuation (e.g. "controls.")
--
-- You might find 'map', 'find', 'isAlpha' and 'filter' useful.
--




lineBreaks::[(String, [String])] -> Int -> [String] -> [([String], [String])]
lineBreaks enHyp y x | (((length (concat x))+((length x)-1)) - y)>0 = init (finalcheck (filterhyp (splitlast (joinlist2 sending (listmaker sending)) y)))
                    | otherwise = [(x,[])]
    where sending = puncchecker x ( findmap findx enHyp ((ischeck x)))


joinlist2::[[String]]->[[String]]->[[String]]
joinlist2 x []=[]
joinlist2 x (y:ys) = ((((concat (joinlist x y)):(joinlist2 x ys))))

joinlist:: [[String]]->[String]->[[String]]
joinlist [] y=[]
joinlist (x:xs) y | (length x) == 1 = x:(joinlist xs y) 
                  | otherwise = y:(joinlist xs y) 


splitlast::[[String]] -> Int -> [([String], [String])]
splitlast [] _=[([],[])]
splitlast (x:xs) y= (splitLine x y): (splitlast xs y)


ischeck::[String] -> [[String]]
ischeck []=[]
ischeck (x:xs)= [(isAlphaa x)] : (ischeck xs)

isAlphaa::String -> String
isAlphaa []=[]
isAlphaa (x:xs) | ((isAlpha x)==True)  = x:isAlphaa(xs)
                | otherwise=[]



findmap::([(String,[String])] -> String -> [String])->[(String,[String])] -> [[String]] -> [[String]]
findmap findx enHyp (x:xs)=
    let found= findx enHyp ((\[b]->b)x)
    in if (length found)==1 
      then [found]++(findmap findx enHyp xs)
        else [found]++xs


findx::[(String,[String])] -> String -> [String]
findx enHyp x | length  (filter (\(y,_) -> y == x) enHyp) ==1 = (\[(_,b)]->b)(filter (\(y,_) -> y == x) enHyp)
              | otherwise=[x]


puncchecker::[String]->[[String]] -> [[String]]
puncchecker [] []=[]
puncchecker (x:xs) (y:ys) | (length x) == (length (concat y)) && (length y)==1 = y:puncchecker xs ys
                          | (length x) /= (length (concat y)) && (length y)==1 = [x]:puncchecker xs ys
                          | (length x) /= (length (concat y)) && (length y)>1= ((init y)++[((last y)++".")]):puncchecker xs ys
                          | (length x) == (length (concat y)) && (length y)>1= y:puncchecker xs ys





combi2:: a -> [a] -> Int ->[a]
combi2 x ys     1 = x:ys
combi2 x (y:ys) z = y:combi2 x ys (z-1)

combi::[String]-> Int -> [[String]]
combi x 1=[]
combi x y | y>1 = (combi2 "-" x y):(combi x (y-1))

listmaker::[[String]] -> [[String]]
listmaker [] =[]
listmaker (x:xs) | ((length x) > 1) =  [(concat x)]:(comblist( combi x ((length x))))
                | otherwise = listmaker xs



comblist::[[String]]->[[String]]
comblist []=[]
comblist (x:xs) = conlist (splitAt ((findd x "-")+1) x):(comblist xs)


findd::[String] -> String -> Int
findd enHyp x =  (((\ (Just (a)) -> a) (findIndex (==x) enHyp)))

conlist::([String],[String])->[String]
conlist (x,y)=[concat x]++[concat y]


filterhyp::[([String],[String])]->[([String],[String])]
filterhyp []=[]
filterhyp (x:xs) | ('-' `elem`  (concat ((\(a,b)->a)x))) &&  (last (concat ((\(a,b)->a)x)) == '-') = x:filterhyp xs
                 | ('-' `elem`  (concat ((\(a,b)->a)x))) &&  (last (concat ((\(a,b)->a)x)) /= '-')= filterhyp xs
                 | otherwise = x:filterhyp xs

finalcheck::[([String], [String])] -> [([String], [String])] 
finalcheck ((x,y):xs)=init ((x,y):(checkerhelper xs (length x)))


checkerhelper::[([String], [String])]->Int -> [([String], [String])] 
checkerhelper [] _=[([],[])]
checkerhelper ((x,y):xs) z | (length x)/=z = (x,y):(checkerhelper xs z)
                     | otherwise = (checkerhelper xs z) 











-- ========================================================= PART 3 ========================================================= --


--
-- Define a function that inserts a given number of blanks (spaces) into a list of strings and outputs a list of all possible
-- insertions. Only insert blanks between strings and not at the beginning or end of the list (if there are less than two
-- strings in the list then return nothing). Remove duplicate lists from the output.
-- The function should take the number of blanks and the the list of strings as input and return a lists of strings.
--
-- Example:
--    blankInsertions 2 ["A", "creative", "man"]   ==>   [["A", " ", " ", "creative", "man"], ["A", " ", "creative", " ", "man"], ["A", "creative", " ", " ", "man"]]
--
-- Use let/in/where to make the code readable
--

combin:: a -> [a] -> Int ->[a]
combin x ys     1 = x:ys
combin x (y:ys) z = y:combin x ys (z-1)

combin1::[String]-> Int -> [[String]]
combin1 x 1=[]
combin1 x y | y>1 = (combin " " x y):(combin1 x (y-1))


combifunc::Int ->[[String]] ->  [[String]]
combifunc 0 []=[]
combifunc _ []=[]
combifunc 0 y = y
combifunc x y |x>0= combifunc (x-1) (removing (map (\z -> (combin1 z (length z))) y))

removing::[[[String]]]->[[String]]
removing []=[]
removing (x:xs) = x ++ removing xs




removedups::[[String]]->[[String]]
removedups []=[]
removedups (x:xs) | x `elem` xs = (removedups xs)
                  | otherwise = x:(removedups xs)



blankInsertions :: Int -> [String] -> [[String]]
blankInsertions x y= removedups (combifunc x (y:[]))







-- ========================================================= PART 4 ========================================================= --


--
-- Define a function to score a list of strings based on four factors:
--
--    blankCost: The cost of introducing each blank in the list
--    blankProxCostfn: The cost of having blanks close to each other
--    blankUnevenCost: The cost of having blanks spread unevenly
--    hypCost: The cost of hyphenating the last word in the list
--
-- The cost of a list of strings is computed simply as the weighted sum of the individual costs. The blankProxCostfn weight equals
-- the length of the list minus the average distance between blanks (0 if there are no blanks). The blankUnevenCost weight is
-- the variance of the distances between blanks.
--
-- The function should take a list of strings and return the line cost as a double
--
-- Example:
--    lineCost ["He", " ", " ", "who", "controls"]
--        ==>   blankCost * 2.0 + blankProxCostfn * (5 - average(1, 0, 2)) + blankUnevenCost * variance(1, 0, 2) + hypCost * 0.0
--        ==>   blankCost * 2.0 + blankProxCostfn * 4.0 + blankUnevenCost * 0.666...
--
-- Use let/in/where to make the code readable
--


---- Do not modify these in the submission ----
blankCost = 1.0
blankProxCost = 1.0
blankUnevenCost = 1.0
hypCost = 1.0
-----------------------------------------------



countSpace::[String]-> Double
countSpace []=0
countSpace (x:xs) | x==" " = 1 + (countSpace xs)
                 | otherwise = countSpace xs

countHyphen::[String]->Double
countHyphen x | last (last x)=='-' =1
              | otherwise =0


blankProxCostfn::[String]->[Float]
blankProxCostfn x| (length (findIndices (==" ") x) >0)=  map realToFrac (blankproxhelper (findIndices (==" ") x) (length x)) 
                | otherwise = [0]


blankproxhelper::[Int] -> Int -> [Int]
blankproxhelper x y= ((\a->[a])(head x))++ (nexthelper x) ++ ((\a->[a])(y-(last x)-1))

nexthelper::[Int] -> [Int]
nexthelper []=[0]
nexthelper (x:xs) | (length xs) >=1 = ((head xs)-x-1):(nexthelper xs) 
                | otherwise = []

average :: [Float] -> Float
average list = (sum list / fromIntegral (length list))



variance::[Float] -> Float
variance x= sum( map (\y-> (y-average(x))^2 ) x) / fromIntegral ((length x))


lineCost :: [String] -> Double
lineCost x = (realToFrac (variance ( blankProxCostfn x))* blankUnevenCost) + (realToFrac (average ( blankProxCostfn x))*blankProxCost) + (hypCost*(countHyphen x)) + (blankCost*(countSpace x))



-- ========================================================= PART 5 ========================================================= --


-- --
-- -- Define a function that returns the best line break in a list of words given a cost function, a hyphenation map and the maximum
-- -- line width (the best line break is the one that minimizes the line cost of the broken list).
-- -- The function should take a cost function, a hyphenation map, the maximum line width and the list of strings to split and return
-- -- a pair of lists of strings as in part 1.
-- --
-- -- Example:
-- --    bestLineBreak lineCost enHyp 12 ["He", "who", "controls"]   ==>   (["He", "who", "cont-"], ["rols"])
-- --
-- -- Use let/in/where to make the code readable
-- --



bestLineBreak :: ([String] -> Double) -> [(String, [String])] -> Int -> [String]-> ([String], [String])
bestLineBreak lineCost enHyp y []=([],[])
bestLineBreak lineCost enHyp y x | ((length (concat x)) + (length x) - 1) < y = (x,[])
bestLineBreak lineCost enHyp y x= 
    let a1=(((\[(a,b)] ->( [a]++[b]) ) (beststring (lineBreaks enHyp y x))))
        a2=(head a1)
    in if (length (head a1))==1
        then ((head a1),concat (tail a1)) 
        else selectone (spaceadd a1 y) (head (spaceadd a1 y))



beststring::[([String], [String])] -> [([String], [String])] 
beststring [(x,[])]=[(x,[])]
beststring (x:xs)= 
    let prevstr=(x:xs)
        newstr= (\[a]->a) (largest (map (\(a,_)-> a) (x:xs)))
    in retstr prevstr newstr

retstr::[([String], [String])]-> [String] ->[([String], [String])]
retstr [] _=[]
retstr (x:xs) y | y == ((\(a,b)->a) x) = [x]
                | otherwise = retstr xs y

largest :: [[String]] -> [[String]]
largest [] = []
largest [y] = [y]
largest (x:y:list)
   | length x < length y = largest (y:list)
   | otherwise = largest (x:list)

spaceadd::[[String]] -> Int -> [([String], [String])]
spaceadd [] _=[]
spaceadd (x:xs) y | (((length x)-1) + (length (concat x))) == y = [(x,(\[a]->a)xs)]
                  | xs /= [] = map (\y->(y,(\[a]->a)xs)) (blankInsertions (y - (length (concat x)) - ((length x)-1)) x)
                  | xs == [] = map (\y->(y,[])) (blankInsertions (y - (length (concat x)) - ((length x)-1)) x)


selectone::[([String], [String])] -> ([String], [String]) -> ([String], [String])
selectone (x:xs) y | (xs)==[] = y
selectone [] y= y
selectone (x:xs) y | (lineCost (((\(a,_)->a))(head xs))) <= (lineCost (((\(a,_)->a)) y)) = (selectone xs (head xs))
                    | otherwise = (selectone xs y)




-- --
-- -- Finally define a function that justifies a given text into a list of lines satisfying a given width constraint.
-- -- The function should take a cost function, hyphenation map, maximum line width, and a text string as input and return a list of
-- -- strings.
-- --
-- -- justifyText lineCost enHyp 15 text1'' should give you the example at the start of the assignment.
-- --
-- -- You might find the words and unwords functions useful.
-- --


bestchecker::([String] -> Double) -> [(String, [String])] -> Int -> [String]-> [[String]]
bestchecker lineCost enHyp y []=[]
bestchecker lineCost enHyp y x | ((length(concat x))+ ((length x)-1)) < y= [x]
bestchecker lineCost enHyp y x =
    let a = bestLineBreak lineCost enHyp y x
        b =  bestchecker lineCost enHyp y ((\(_,b)->b)a)
    in ((\(a,b)->a)a):b



justifyText :: ([String] -> Double) -> [(String, [String])] -> Int -> String -> [String]
justifyText lineCost enHyp y x = init( concat(map (\y->y++["\n"]) (bestchecker lineCost enHyp y (splitOn " " x))))















