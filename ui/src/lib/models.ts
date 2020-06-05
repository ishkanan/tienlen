
export enum Suit {
  Spades = 1,
  Clubs = 2,
  Diamonds = 3,
  Hearts = 4,
}

export interface Card {
  suit: Suit;
  faceValue: number;
  suitRank: number;
  globalRank: number;
}

export interface Player {
  id: string;
  name: string;
  position: number;
  cardsLeft: number;
  isPassed: boolean;
  isTurn: boolean;
  wonLastGame: boolean;
  connected: boolean;
}

export enum GameEventKind {
  Info = 1,
  Error = 2
}

export interface GameEvent {
  kind: GameEventKind;
  message: string;
}
