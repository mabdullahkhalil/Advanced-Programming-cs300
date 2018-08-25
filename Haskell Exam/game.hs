-- CS300 Spring 2018 Exam 1 - 25 February 2018
--
-- Total time is 5 hours and maximum score is 25 points.
-- Parts *must* be done in order. Exact score per part will be assigned later. 
-- If a part is not working, you will *not* be graded for any later parts. 
-- No USBs, mobiles, or electronic devices with you. Printouts are okay.
-- Strictly no talking to anyone during the exam. 
--
-- We have given a main function that uses test cases near the end of file to 
-- tell you how many parts you have successfully done. Note that we will test 
-- it on additional inputs for grading so do not hardcode for the given test 
-- cases. You only need to fill in the code for each part. The error 
-- "Prelude.undefined" means that you have not yet written the code for the 
-- next part in order.-
--
-------------------------------------------------------------------------------
import Data.List (sort)
-- YOU ARE NOT ALLOWED TO DO ANY ADDITIONAL IMPORTS
--
-- You are a computer scientist hired by a game development company to improve 
-- the performance of their multiplayer Roll Playing Game (RPG). You realize 
-- that the game stores its state (the position of all players and objects in 
-- the "game world") in a 2-D character array, or equivalently, an array of 
-- "strings". Here is a sample game state:
-- 
--  Col 0  Col 1   Col 2  Col 3
--  ------------------------------------
-- [[' ',   ' ',   ' ',   ' '],        | (Row 0)
--  [' ',   'W',   'W',   ' '],        | (Row 1)
--  ['E',   ' ',   ' ',   ' '],        | (Row 2)
--  [' ',   ' ',   'P',   ' ']]        | (Row 3)
--
-- The ' ' characters signify empty spaces in the game. The player can walk 
-- through these. 'E' is the enemy sprite (a game object is called a "sprite").
-- 'W' are walls. Sprite positions can be marked using any symbol other than 
-- ' ' which, as mentioned earlier, is for empty spaces in the game world.
--
-- To begin your analysis, you decide to write a couple of helper functions 
-- which you can call every turn to analyze what the game world contains.

-- === Part 1 ===
-- Returns the number of rows in a world
-- FOR THIS PART, YOU CANNOT USE ANY BUILTIN FUNCTION
numRows :: [[a]] -> Int
numRows [] = 0
numRows (x:xs) = 1+ (numRows xs) 

-- === Part 2 ==
-- Given the value that represents an empty location, the number of rows and 
-- the number of columns (in this order), create a new game world
createEmptyWorld :: a -> Int -> Int -> [[a]]
createEmptyWorld _ _ 0=[]
createEmptyWorld a x y=[(replicate x a)]++(createEmptyWorld a x (y-1))

-- === Part 3 ===
-- Returns True if the row contains *only* sprites (or empty spaces) with the
-- symbol given in argument and no other sprites.
-- As an example, in the sample game state, rowContainsOnly called with ' ' 
-- will return True on Row 0, but False on Rows 1 to 3.

-- check for W
rowContainsOnly :: Eq a => a -> [a] -> Bool
rowContainsOnly a x | (all (a==) x) ==True = True
                    | otherwise = False
-- === Part 4 ===
-- Returns True if entire world contains *only* sprites (or empty space) with 
-- the symbol given in argument.
worldContainsOnly :: Eq a => a -> [[a]] -> Bool
worldContainsOnly _ []=True
worldContainsOnly a (x:xs) = (all (a==) x) && (worldContainsOnly a xs)

-- === Part 5 ===
-- Given a world, return a new world with the same dimensions but with only 
-- empty spaces where the value representing empty space is given as argument
clearWorld :: a -> [[a]] -> [[a]]
clearWorld _ []=[]
clearWorld a (x:xs) = [(replicate (length x) a) ]++(clearWorld a xs)

-- === Part 6 ===
-- Returns True if a given world contains *at least one* sprite with the symbol
-- given in first argument
worldContains :: Eq a => a -> [[a]] -> Bool
worldContains _ []=False
worldContains a (x:xs) = (a `elem` x) || (worldContains a xs)

-- == Part 7 ==
-- After analysis, you realize that it is not the processing time on the game 
-- servers that is causing delay, it is the *network* which is the bottleneck. 
-- The company has a proprietary network which its users subscribe to. The 
-- network is very low latency but has very little bandwidth. The game sends 
-- its game state to the server after every turn, but the game state in its 
-- current format is large enough to cause a lot of delay. What is needed is to 
-- find ways to compress the game state. You figure out two ways to compress 
-- the game state, which you plan to use in two different cases.
--
-- First arpproach: The game state is composed almost entirely of empty spaces 
-- and the number of sprites is very small. For this sutation, you represented 
-- every sprite (excluding empty spaces) with a tuple containing its symbol, 
-- its row in the game world, and its column in the game world. The compressed 
-- game state should look like this: 
-- [(s1, r1, c1), (s2, r2, c2), ... ]
--
-- Your task is to take the symbol to represent empty spaces (e.g. ' '), row 
-- number, column number and the actual row (char array / string) in this order 
-- and return the compressed form of the row as an array of tuples.
easyCompressRow :: Eq a => a -> Int -> Int -> [a] -> [(a, Int, Int)]
easyCompressRow _ _ _ []=[]
easyCompressRow x r c (y:ys) | x/=y = [(y,r,m)] ++ (easyCompressRow x r (c) ys)
                             | otherwise = (easyCompressRow x r c ys)
                             where m= (c-(length (y:ys)))

-- == Part 8 ==
-- Takes in the symbol representing empty spaces and a game world and returns 
-- its compressed version as an array of tuples
rowfunc::Eq a => a -> Int -> [a] -> [(a, Int, Int)]
rowfunc a r x= easyCompressRow a r (length x) x 

easyCompressWorld :: Eq a => a -> [[a]] -> [(a, Int, Int)]
easyCompressWorld _ [] =[]
easyCompressWorld a b = 
    let row = (rowfunc a ((length b)-1) (last b))
        world = row ++ (easyCompressWorld a (init b))
    in reverse(world) 

-- == Part 9 ==
-- Given a sprite symbol, the desired column, and a row of sprites (in that
-- order), return a new row with the new sprite inserted in the desired column
insertSpriteInRow :: a -> Int -> [a] -> [a]
insertSpriteInRow a c l = ((take (c) l)++[a] ++(drop (c+1) l))

-- == Part 10 == 
-- Insert the given sprite and the given row and column in the given wrold (in 
-- that order) and return a new world with the desired sprite added
insertSpriteInWorld :: a -> Int -> Int -> [[a]] -> [[a]]
insertSpriteInWorld s r c l = ((take (r) l)++[(insertSpriteInRow s c (l !! r))] ++ (drop (r+1) l))

-- == Part 11 ==
-- Given the empty symbol, the rows and columns, and the compressed world, 
-- return the corresponding decompressed world

addingfunc::[(a, Int, Int)] -> [[a]] ->[[a]]
addingfunc [] _=[]
addingfunc x world | length x==1 = (\[(m,n,o)] -> (insertSpriteInWorld m n o world)) x
addingfunc (x:xs) world =
    let one= (\(m,n,o) -> (insertSpriteInWorld m n o world)) x
        two = addingfunc xs one
    in two



easyDecompressWorld :: Eq a => a -> Int -> Int -> [(a, Int, Int)] -> [[a]]
easyDecompressWorld a r c x=
     let world= (createEmptyWorld a r c)
         adding= addingfunc x world
     in adding


-- == Part 12 ==
-- Second approach: When the game world is filled with sprites, the above 
-- compression performs even worse than the original implementation. So you 
-- decide to create something better for this case. You realize the game only 
-- has 3 sprites in addition to empty spaces. These are Walls (W), Enemies (E) 
-- and the Player (P). You also realize that the order of sprites in a row does 
-- not matter. What matters is that sprites in one row should remain in the 
-- same row when they reach the server, even if their order within that row 
-- changes. You decide to represent every sprite with a prime number and every 
-- row with a product of those primes. You make the following algorithm:
-- Begin with 1
-- For each ' ' in the row, multiply by 2.
-- For each W in the row, multiply by 3.
-- For each E in the row, multiply by 5.
-- For each P in the row, multiply by 7.
-- Return one integer for each row.
-- e.g [' ', ' ', 'W', 'P'] becomes 2 * 2 * 3 * 7 which is 84
-- It does not matter what you return for invalid inputs
charToPrime :: [Char] -> Int
charToPrime []=1
charToPrime (x:xs) | x==' '= 2*(charToPrime xs)
                   | x=='W'= 3*(charToPrime xs)
                   | x=='E'= 5*(charToPrime xs)
                   | x=='P'= 7 *(charToPrime xs)
-- == Part 13 ==
-- Since each sprite is represented by a prime, the product representing the 
-- row has a unique factorizing into the primes allowing us to figure out what
-- the original sprites in the row were. You can use integer modulus operator 
-- (mod) and integer division operator (div). For example "84 `mod` 2==0" which 
-- means 2 is a prime factor and hence there was a ' ' in the row. We can use 
-- div to find leftover product e.g. "84 `div` 2==42" which means 42 is the 
-- product representing the remaining sprites. It is again divisble by 2 which
-- means there was another ' ' in the row and so on. It does not matter what 
-- you return for invalid inputs.
primeToChar :: Int -> [Char]
primeToChar i | i==1=[]
              | i `mod` 2 == 0 = " " ++ (primeToChar (i `div` 2))
              | i `mod` 3 == 0 = "W" ++ (primeToChar (i `div` 3))
              | i `mod` 5 == 0 = "E" ++ (primeToChar (i `div` 5))
              | i `mod` 7 == 0 = "P" ++ (primeToChar (i `div` 7))
-- == Part 14 ==
-- The class below represents a type of sprite whose lists can be converted 
-- back and forth from integers. You have to make the Char type a member of the
-- Pr typeclass by implementing the required functions.
class Pr a where
    toPrime :: [a] -> Int
    fromPrime :: Int -> [a]

instance Pr Char where
    toPrime = charToPrime
    fromPrime = primeToChar

mytest::Int->[Char]
mytest x= fromPrime x


-- == Part 15 ==
-- Returns the list of compressed integers repr senting the entire game world 
betterCompressWorld :: Pr a => [[a]] -> [Int]
betterCompressWorld world = map (toPrime) world

-- == Part 16 ==
-- Decompress the list of primes to the game world
betterDecompressWorld :: Pr a => [Int] -> [[a]]
betterDecompressWorld world = map (fromPrime) world

-- == Part 17 ==
-- Recall the Either data type
--   data Either a b = Left a | Right b
-- You have to decompress the world when it is compressed in either of the two
-- approaches discussedi. You are given the empty value, the umber of rows and
-- columns and the compressed world represented as an Either value
decompressWorld :: (Eq a, Pr a) => a -> Int -> Int -> Either [(a,Int,Int)] [Int] -> [[a]]
decompressWorld _ _ _ (Right x) = betterDecompressWorld x
decompressWorld a r c (Left x) = easyDecompressWorld a r c x
-- == Part 18 ==
-- Intelligent compression decision:
-- The first approach requires three integers to represent every non-empty 
-- sprite while the second approach requires one integer for every row.
-- Compress using the approach that requires fewer total integers. Count each
-- tuple of the tuple approach as two integers. It does not matter what you
-- choose if the number of integers is equal for a given world.
compressWorld :: (Eq a, Pr a) => a -> [[a]] -> Either [(a,Int,Int)] [Int]
compressWorld sp world= 
    let a= easyCompressWorld sp world 
        b= betterCompressWorld world
    in if ((length a ) * 2) > length b 
        then Right b 
        else Left a


-- == Part 19 ==
-- We realize that using Char to represent Sprite was not a great idea. Luckily
-- we did not hard-code anything for Char. So we create a data type called 
-- Sprite. You may ignore the "deriving" part 
data Sprite = Player | Enemy | Wall deriving (Show, Eq)
-- To signify an Empty space we will use Maybe Sprite. Your task is to make 
-- "Maybe Sprite" an instance of Pr such that all previous functions start to 
-- work for it.

-- instance  Maybe Sprite where
--     Just Player =  charToPrime

-- == Part 20 ==
-- Finally we want to convert a world represented using Chars to our newer
-- format of representing worlds using Maybe Sprites. But since both can be 
-- represented using primes, we come up with a more abstract function that can
-- convert between any two types of sprites that are members of the Pr 
-- typeclass.  When testing this function, you have to call it from a context 
-- where the return type can be deduced, otherwise the compiler will complain 
-- that it cannot determine what b is. Since you are on the last part, this is 
-- the only hint you will get :)
convert :: (Pr a, Pr b) => [[a]] -> [[b]]
convert _ = undefined
