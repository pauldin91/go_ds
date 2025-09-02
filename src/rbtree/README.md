# rightRotation(x) --> lean to right subtree (clockwise rotation)

## case 1

        x                 y
       / \               / \
      y   ?   --->      a   x
     / \		         / \
    a   b		        b   ?

## case 2

        x                 y
       / \               / \
      y   b   --->      a   x
     / \		         / \
    a   ?		        ?   b

# leftRotation(x) --> lean to left subtree (counterclockwise rotation)

## case 1

      x                    y
     / \                  / \
    ?   y    --->        x   b
       / \              / \
      a   b            ?   a

## case 2

      x                    y
     / \                  / \
    a   y    --->        x   b
       / \              / \
      ?   b            a   ?
