import type { Question } from './question';

export type Diary = {
  id: number;
  note: string;
  userId: number;
  question: Question;
  createdAt: Date;
};
