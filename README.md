## A Secret Role Card Game (or... it will be)

This project came to be as an opportunity for me to work on a solo game and to learn Go (which I've grown to like... mostly. I wish there were enums lol).
I technically first came up with this idea as a game for my university mobile game development class, but we ended up moving forward with a different idea that was maybe better (and that turned into _Apathia_).
I also figured I need more public examples of my code, so here you are! The code is a bit messy (WIP!!!), but I hope it demonstrates my abilities to tackle difficult problems and write complex, robust systems.

As of now, this is very much a work in progress (when writing this, I am just finishing off the main mechanical loop for the first real prototype).

## Main Rules
(these may not be fully up to date)
* #### Secret Roles: Everyone gets a secret role at the beginning of the game that effects how they play (win conditions?) and their deck(?)
	* ~ 3-5 players
	* Roles have unique passives
* ### Turn Phases
	* ##### Card Passing:
		* Players can pass any number of cards in their hand to any number of players. These passes are only locked in (and their effects applied) at the end of this phase. The player with the most cards in hand at the end of this phase gets a failure point (or a damage... idk).
	* ##### Playing: 
		* In clockwise order (starts at least VP or in rotating spot), players get to do 2 actions. They HAVE TO DO BOTH (if possible).
		* 1) Play a card they kept in their hand
		* 2) Play a card that was passed to them
		* Only the played cards are revealed; the rest are secret
		* Not sure if they need to be done in this order or if the order can be up to the player (prob without public info as to which was which)
		* Maybe some other shit too (e.g., legendary actions), idk
* ### Other rules
	* ~ Card Drawing:
		* Deck is ~30 cards
		* Beginning of each turn, ~5 cards in the deck randomly upgrade and each player draws 5 cards from the deck
		* End of turn, all cards in play and in hands are shuffled together and placed at the bottom of the deck
	* Cards:
		* Cards come in 3 levels--bronze, silver, gold (lvl 1, 2, 3)
		* 2 Types of cards:
			* Normal (just play them, no target, effect just happens)
			* Targeted (comes with a target stapled on the card; that being Left or Right)
		* Could allow cards to be passed to any player, or maybe just L, R, across (idk)
		* Maybe some roles add certain cards to the deck?
	* Strength:
		* Your strength = your cards in hand (+ some modifiers).
		* = cards that you held + cards that were passed to you (-1 or -2 if you are actively playing a card)

### Card Resolution
* Cards are composed of several (trigger, effect) pairs we will call Abilities
	* triggers are Resolve, NextPlay, Draw, etc.
* When a card is played, its resolve ability goes on the stack
	* A resolve trigger automatically queues its other triggers at the end of its resolution
	* The stack is composed of abilities, not cards
* A card is taken off the stack after it resolves (and the stack is empty) and is put into the waiting zone (which is then shuffled and put on the bottom of the deck)
	* really, the card itself is in its own zone while on the stack

----------------------------------

Some magic the gathering card arts are used as placeholder art. This is intended to be temporary, and thus will be removed as soon as the game is at all playable. Until that moment, the art is used under WOTC fan content policy: https://company.wizards.com/en/legal/fancontentpolicy.
If there is any issue regarding copyright or any other matter, do not hesitate to reach out to me.
