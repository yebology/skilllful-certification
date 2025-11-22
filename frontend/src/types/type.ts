export type Class = {
  id: number;
  name: string;
  category: string;
};

export type ClassDetail = {
  id: number;
  name: string;
  category: string;
  description: string;
  instructor: string;
};

export type Participant = {
  id: number;
  name: string;
  email: string;
  gender: string;
};

export type ParticipantDetail = {
  id: number;
  name: string;
  email: string;
  gender: string;
  phoneNumber: string;
  birthDate: string;
};