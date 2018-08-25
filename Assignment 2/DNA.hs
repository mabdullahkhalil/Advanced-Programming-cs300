-- ---------------------------------------------------------------------
-- DNA Analysis 
-- CS300 Spring 2018
-- Due: 24 Feb 2018 @9pm
-- ------------------------------------Assignment 2------------------------------------
--
-- >>> YOU ARE NOT ALLOWED TO IMPORT ANY LIBRARY
-- Functions available without import are okay
-- Making new helper functions is okay
--
-- ---------------------------------------------------------------------
--
-- DNA can be thought of as a sequence of nucleotides. Each nucleotide is 
-- adenine, cytosine, guanine, or thymine. These are abbreviated as A, C, 
-- G, and T.
--



type DNA = [Char]
type RNA = [Char]
type Codon = [Char]
type AminoAcid = Maybe String


-- ------------------------------------------------------------------------
-- 				PART 1
-- ------------------------------------------------------------------------				

-- We want to calculate how alike are two DNA strands. We will try to 
-- align the DNA strands. An aligned nucleotide gets 3 points, a misaligned
-- gets 2 points, and inserting a gap in one of the strands gets 1 point. 
-- Since we are not sure how the next two characters should be aligned, we
-- try three approaches and pick the one that gets us the maximum score.
-- 1) Align or misalign the next nucleotide from both strands
-- 2) Align the next nucleotide from first strand with a gap in the second     
-- 3) Align the next nucleotide from second strand with a gap in the first    
-- In all three cases, we calculate score of leftover strands from recursive 
-- call and add the appropriate penalty.                                    

score :: DNA -> DNA -> Int
score [] []=0
score [] (y:ys) = 1+ (score [] ys) 
score (x:xs) [] = 1+ (score xs []) 
score (x:xs) (y:ys) | x==y = 3+ (score xs ys)
                    | otherwise = (max misAlign (max sndSpace fstSpace))
                    where
                        misAlign = 2 + (score xs ys)
                        sndSpace = 1 + (score xs (y:ys))
                        fstSpace = 1 + (score (x:xs) ys)


-- -------------------------------------------------------------------------
--				PART 2
-- -------------------------------------------------------------------------
-- Write a function that takes a list of DNA strands and returns a DNA tree. 
-- For each DNA strand, make a separate node with zero score 
-- in the Int field. Then keep merging the trees. The merging policy is:
-- 	1) Merge two trees with highest score. Merging is done by making new
--	node with the smaller DNA (min), and the two trees as subtrees of this
--	tree
--	2) Goto step 1 :)
--

data DNATree = Node DNA Int DNATree DNATree | Nil deriving (Ord, Show, Eq)

findMaxhelper::[DNATree]->[(DNATree,DNATree,Int)]
findMaxhelper x | length x ==1 =[]
findMaxhelper (x:xs) | ((length xs) >=1)=(findMax x xs)++(findMaxhelper xs)


findMax::DNATree -> [DNATree]->[(DNATree,DNATree,Int)]
findMax _ []=[]
findMax x (y:ys)=
    let xx= (\(Node a b c d)-> a)x
        yy= (\(Node a b c d)-> a)y
    in [(x,y,(score xx yy))]++findMax x ys


maxiDNAmatch::[(DNATree,DNATree,Int)]->[(DNATree,DNATree,Int)]
maxiDNAmatch [] = [(Nil,Nil,0)]
maxiDNAmatch [y] = [y]
maxiDNAmatch (x:y:xs)=
    let xx= (\(a,b,c)->c) x
        yy = (\(a,b,c)->c) y
    in if xx<yy
        then  maxiDNAmatch (y:xs)
        else if xx==yy
        then x:maxiDNAmatch (y:xs)
        else maxiDNAmatch (x:xs)


makeTree::(DNATree,DNATree,Int)->DNATree
makeTree (x,y,z) =
    let xx = (\(Node a b c d)-> a) x
        yy = (\(Node a b c d)-> a) y
    in if length xx <= length yy 
        then (Node xx z x y)
        else (Node yy z y x)


findmaxDNA::[(DNATree,DNATree,Int)]->(DNATree,DNATree,Int)
findmaxDNA []=(Nil,Nil,0)
findmaxDNA x | length x ==1 = (\[a]->a)x
findmaxDNA (x:y:xs) =
    let fstx = (\(Node a b c d)-> a) ((\(a,b,c)->a) x)  
        sndx = (\(Node a b c d)-> a) ((\(a,b,c)->b) x)  
        fsty = (\(Node a b c d)-> a) ((\(a,b,c)->a) y)  
        sndy = (\(Node a b c d)-> a) ((\(a,b,c)->b) y)  
        minx = minimum[fstx ,sndx] 
        miny = minimum[fsty,sndy]
        maxboth= maximum [minx,miny]
    in if fstx==maxboth || sndx==maxboth
        then  (findmaxDNA (x:xs))
        else  (findmaxDNA (y:xs))



assistDNA::[DNATree]->DNATree
assistDNA trees | (length trees ==1) = (\[a]->a)trees
assistDNA trees = 
    let maxfind = (findMaxhelper trees)
        maxDNA = findmaxDNA (maxiDNAmatch (maxfind))
        treemade= makeTree maxDNA
        removed= removeAlready trees treemade
        final= assistDNA removed
    in final



makeDNATree :: [DNA] -> DNATree
makeDNATree x =
    let trees = (map (\y-> Node y 0 Nil Nil) x)
    in (assistDNA trees)

removeAlready::[DNATree]->DNATree->[DNATree]
removeAlready [] y=[y]
removeAlready (x:xs) y | aa==yy || bb==aa = (removeAlready xs y)
                       | otherwise = x:(removeAlready xs y)
                where aa= (\(Node a b c d)-> a) x -- Node aa b c d = x
                      bb=(\(Node a b c d)-> a)  y -- Node 
                      yy=(\(Node a b c d)-> a) ((\(Node a b c d)-> d) y)




-- -------------------------------------------------------------------------
--				PART 3
-- -------------------------------------------------------------------------

-- Even you would have realized it is hard to debug and figure out the tree
-- in the form in which it currently is displayed. Lets try to neatly print 
-- the DNATree. Each internal node should show the 
-- match score while leaves should show the DNA strand. In case the DNA strand 
-- is more than 10 characters, show only the first seven followed by "..." 
-- The tree should show like this for an evolution tree of
-- ["AACCTTGG","ACTGCATG", "ACTACACC", "ATATTATA"]
--
-- 20
-- +---ATATTATA
-- +---21
--     +---21
--     |   +---ACTGCATG
--     |   +---ACTACACC
--     +---AACCTTGG
--
-- Make helper functions as needed. It is a bit tricky to get it right. One
-- hint is to pass two extra string, one showing what to prepend to next 
-- level e.g. "+---" and another to prepend to level further deep e.g. "|   "


draw :: DNATree -> [Char]
draw x | x==Nil = "aaa"
       | (((\(Node _ _ a _)->a) x) == Nil) && (((\(Node _ _ _ a)->a) x)==Nil) =( (\(Node a _ _ _)->a) x)++ "\n|+---"
draw x=
    let score= (show ((\(Node _ a _ _)->a) x))
        left= draw ((\(Node _ _ a _)->a) x)
        right= draw ((\(Node _ _ _ a)->a) x)
    in if (((\(Node _ a _ _)->a) x)) >0 
       then (concat (pipeadd (words ( makeStr (score ++"\n|+---"++ left++right))) 0))
       else ""


makeStr::[Char]->[Char]
makeStr [] =[]
makeStr (x:xs) | x `elem` ['\n','|','+','-'] = " "++makeStr xs
                 | otherwise = [x] ++ makeStr xs


pipeadd::[String]->Int->[String]
pipeadd [] _=[]
pipeadd (x:xs) y | (all(`elem` ['A'..'Z']) x) && y==0 = ([concat ["+","-","-","-"]]) ++ [x] ++ ["\n"] ++ (pipeadd xs y)
       | (all(`elem` ['A'..'Z']) x) && y>=1 = (replicate (y-1) "|   ")++([concat ["+","-","-","-"]]) ++ [x] ++ ["\n"] ++ (pipeadd xs y)
       |  y==0 = [x]++["\n"]++(pipeadd xs (y+1))
       |  y>=1 = (replicate (y-1) "|   ")++["+---"]++[x]++["\n"]++(pipeadd xs (y+1))


-- ---------------------------------------------------------------------------
--				PART 4
-- ---------------------------------------------------------------------------
--
--
-- Our score function is inefficient due to repeated calls for the same 
-- suffixes. Lets make a dictionary to remember previous results. First you
-- will consider the dictionary as a list of tuples and write a lookup
-- function. Return Nothing if the element is not found. Also write the 
-- insert function. You can assume that the key is not already there.

type Dict a b = [(a,b)]

lookupDict :: (Eq a) => a -> Dict a b -> Maybe b
lookupDict x [] = Nothing
lookupDict x (y:ys) | (x==((\(a,b)->a) y)) = Just ((\(a,b)->b) y)
                    | otherwise = lookupDict x ys



insertDict :: (Eq a) => a -> b -> (Dict a b)-> (Dict a b)
insertDict x y z = z ++ [(x,y)]


-- We will improve the score function to also return the alignment along
-- with the score. The aligned DNA strands will have gaps inserted. You
-- can represent a gap with "-". You will need multiple let expressions 
-- to destructure the tuples returned by recursive calls.

alignment1 :: DNA -> DNA -> [((Char, Char), Int)]
alignment1 [] []=[]
alignment1 [] (y:ys) = [(('-',y),1)] ++ (alignment1 [] ys) 
alignment1 (x:xs) [] = [((x,'-'),1)] ++ (alignment1 xs []) 
alignment1 (x:xs) (y:ys) | x==y = [((x,y),3)] ++ (alignment1 xs ys)
                       | (x/=y) && (misAlign == (max misAlign (max sndSpace fstSpace))) = [((x,y),2)] ++ (alignment1 xs ys)
                       | (x/=y) && (sndSpace == (max misAlign (max sndSpace fstSpace))) = [((x,'-'),1)] ++ (alignment1 xs (y:ys))
                       | (x/=y) && (fstSpace == (max misAlign (max sndSpace fstSpace))) = [(('-',y),1)] ++ (alignment1 (x:xs) ys)
                    where
                        misAlign = 2 + (sum (map (\(a,b)->b) (alignment1 xs ys)))
                        sndSpace = 1 + (sum (map (\(a,b)->b) (alignment1 xs (y:ys))))
                        fstSpace = 1 + (sum (map (\(a,b)->b) (alignment1 (x:xs) ys)))


alignment :: String -> String -> ((String, String), Int)
alignment x y =
    let all=alignment1 x y
        allA= map (\((a,_),_)->(a)) all
        allB= map (\((_,b),_)->(b)) all
        allC= sum (map (\((_,_),a)->(a)) all)
    in (((allA),(allB)),allC)


-- We will now pass a dictionary to remember previously calculated scores 
-- and return the updated dictionary along with the result. Use let 
-- expressions like the last part and pass the dictionary from each call
-- to the next. Also write logic to skip the entire calculation if the 
-- score is found in the dictionary. You need just one call to insert.
type ScoreDict = Dict (DNA,DNA) Int

scoreMemo :: (DNA,DNA) -> ScoreDict -> (ScoreDict,Int)
scoreMemo (x,y) z | lookup == Nothing = ((insertDict (x,y) (score x y) z),(score x y))
                  | otherwise =  (z , ((\(Just a)->a) lookup))
                where lookup = lookupDict (x,y) z


-- In this part, we will use an alternate representation for the 
-- dictionary and rewrite the scoreMemo function using this new format.
-- The dictionary will be just the lookup function so the dictionary 
-- can be invoked as a function to lookup an element. To insert an
-- element you return a new function that checks for the inserted
-- element and returns the old dictionary otherwise. You will have to
-- think a bit on how this will work. An empty dictionary in this 
-- format is (\_->Nothing)

type Dict2 a b = a->Maybe b
-- insertDict2 :: (Eq a) => a -> b -> (Dict2 a b)-> (Dict2 a b)
-- insertDict2 x y f = f x y

type ScoreDict2 = Dict2 (DNA,DNA) Int

-- scoreMemo2 :: (DNA,DNA) -> ((DNA, DNA) -> Maybe Int) -> (ScoreDict2,Int)
-- scoreMemo2 (x, y) f = f (x, y) == Nothing = undefined

-- map :: (a -> b) -> [a] -> [b]


-- ---------------------------------------------------------------------------
-- 				PART 5
-- ---------------------------------------------------------------------------

-- Now, we will try to find the mutationDistance between two DNA sequences.
-- You have to calculate the number of mutations it takes to convert one 
-- (start sequence) to (end sequence). You will also be given a bank of 
-- sequences. However, there are a couple of constraints, these are as follows:

-- 1) The DNA sequences are of length 8
-- 2) For a sequence to be a part of the mutation distance, it must contain 
-- "all but one" of the neuclotide bases as its preceding sequence in the same 
-- order AND be present in the bank of valid sequences
-- 'AATTGGCC' -> 'AATTGGCA' is valid only if 'AATTGGCA' is present in the bank
-- 3) Assume that the bank will contain valid sequences and the start sequence
-- may or may not be a part of the bank.
-- 4) Return -1 if a mutation is not possible

	
-- mutationDistance "AATTGGCC" "TTTTGGCA" ["AATTGGAC", "TTTTGGCA", "AAATGGCC", "TATTGGCC", "TTTTGGCC"] == 3
-- mutationDistance "AAAAAAAA" "AAAAAATT" ["AAAAAAAA", "AAAAAAAT", "AAAAAATT", "AAAAATTT"] == 2



checkDist:: DNA -> DNA ->DNA -> Int -> [DNA]
checkDist _ _ _ 8 = []
checkDist (x:xs) (y:ys) a z | x/=y  = m ++ (checkDist xs ys ((\[a]->a)m) (z+1))
                            | x==y = (checkDist xs ys a (z+1))
                            where m=[(take z a)++[y]++xs] 


countdist::[DNA]-> [DNA]->Int
countdist [] _ =0
countdist (z:zs) y | (z `elem` y) = 1+(countdist zs y) 
                   | otherwise = -1 +(countdist zs y) 



mutationDistance :: DNA -> DNA -> [DNA] -> Int
mutationDistance x y z =  countdist (checkDist x y x 0) z




-- ---------------------------------------------------------------------------
-- 				PART 6
-- ---------------------------------------------------------------------------
--
-- Now, we will write a function to transcribe DNA to RNA. 
-- The difference between DNA and RNA is of just one base i.e.
-- instead of Thymine it contains Uracil. (U)
--
transcribeDNA :: DNA -> RNA
transcribeDNA []=[]
transcribeDNA (x:xs) | x=='T' = 'U':(transcribeDNA xs)
                     | otherwise= x:(transcribeDNA xs)





-- Next, we will translate RNA into proteins. A codon is a group of 3 neuclotides 
-- and forms an aminoacid. A protein is made up of various amino acids bonded 
-- together. Translation starts at a START codon and ends at a STOP codon. The most
-- common start codon is AUG and the three STOP codons are UAA, UAG and UGA.
-- makeAminoAcid should return Nothing in case of a STOP codon.
-- Your translateRNA function should return a list of proteins present in the input
-- sequence. 
-- Please note that the return type of translateRNA is [String], you should convert
-- the abstract type into a concrete one.
-- You might wanna use the RNA codon table from 
-- https://www.news-medical.net/life-sciences/RNA-Codons-and-DNA-Codons.aspx
-- 
--

makeAminoAcid :: Codon -> AminoAcid
makeAminoAcid x |x `elem` ["AUU","AUC","AUA"]= Just "i"
                |x `elem` ["CUU","CUC","CUA","CUG","U","UA","UUG"]= Just "l"
                |x `elem` ["GUU","GUC","GUA","GUG"]= Just "v"
                |x `elem` ["UUU","UUC"]= Just "f"
                |x `elem` ["AUG"]= Just "m"
                |x `elem` ["UGU","UGC"]= Just "c"
                |x `elem` ["GCU","GCC","GCA","GCG"]= Just "a"
                |x `elem` ["GGU","GGC","GGA","GGG"]= Just "g"
                |x `elem` ["CCU","CCC","CCA","CCG"]= Just "p"
                |x `elem` ["ACU","ACC","ACA","ACG"]= Just "t"
                |x `elem` ["UCU","UCC","UCA","UCG","AGU","AGC"]= Just "s"
                |x `elem` ["UAU","UAC"]= Just "y"
                |x `elem` ["UGG"]= Just "w"
                |x `elem` ["CAA","CAG"]= Just "q"
                |x `elem` ["AAU","AAC"]= Just "n"
                |x `elem` ["CAU","CAC"]= Just "h"
                |x `elem` ["GAA","GAG"]= Just "e"
                |x `elem` ["GAU","GAC"]= Just "d"
                |x `elem` ["AAA","AAG"]= Just "k"
                |x `elem` ["CGU","CGC","CGA","CGG","AGA","AGG"]= Just "r"
                |x `elem` ["UAA","UAG","UAG"]=Nothing
                | otherwise = Just "-"



translate :: RNA -> [String]
translate []=[]
translate x=
    let aminoacidd=makeAminoAcid(take 3 x)
        noNothing=((\(Just a)->[a])aminoacidd ++ (translate (drop 3 x)))
        withNothing = (["-"]++ (translate (drop 3 x)))
    in if aminoacidd /= Nothing
        then noNothing
        else withNothing

strmake::String -> [String]
strmake [] =[]
strmake x | length x ==1 =[]
strmake x = [(takeWhile (/='-') (dropWhile (/='m') x))] ++ (strmake (dropWhile (/='-') (tail x)))

translateRNA :: RNA -> [String]
translateRNA x = strmake (concat (translate x))

