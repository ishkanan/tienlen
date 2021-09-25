import { Card, Player } from './models';

export interface Message {
  kind: string;
  data: string;
}

export interface JoinGameRequest {
  playerName: string;
}

export interface PlayerJoinedResponse {
  player: Player;
}

export interface PlayerDisconnectedResponse {
  player: Player;
}

export interface StartGameRequest {}

export interface GameStartedResponse {
  player: Player;
}

export interface GamePausedResponse {}

export interface GameResumedResponse {}

export interface TurnPassRequest {}

export interface TurnPassedResponse {
  player: Player;
}

export interface RoundWonResponse {
  player: Player;
}

export interface TurnPlayRequest {
  cards: number[];
}

export interface TurnPlayedResponse {
  player: Player;
  cards: Card[];
}

export interface PlayerPlacedResponse {
  player: Player;
  place: number;
}

export interface GameWonResponse {
  player: Player;
}

export interface ResetGameRequest {}

export interface GameResetResponse {
  player: Player;
}

export enum GameState {
  InLobby = 1,
  Running = 2,
  Paused = 3,
}

export interface GameStateRefreshResponse {
  opponents: Player[];
  self: Player;
  selfHand: Card[];
  gameState: GameState;
  lastPlayed: Card[];
  firstRound: boolean;
  newRound: boolean;
  winPlaces: Player[];
}

export enum ErrorKind {
  LobbyNotReady = 1,
  NotAuthorised = 2,
  OutOfTurn = 3,
  MustPlay = 4,
  InvalidCards = 5,
  InvalidPattern = 6,
  CardsNotBetter = 7,
  MustPlayLowest = 8,
  NameTaken = 9,
  GameFull = 10,
}

export interface ErrorResponse {
  kind: ErrorKind;
}
