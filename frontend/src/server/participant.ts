import axios from "axios";
import { API_URL } from "../env/env";
import type { Participant, ParticipantDetail } from "../types/type";

export async function getAllParticipant(): Promise<Participant[]> {
  try {
    const res = await axios.get(API_URL + "/participants");
    return res.data.data as Participant[];
  } catch (error) {
    console.error(error);
    return [];
  }
}

export async function getParticipantDetail(
  id: number
): Promise<ParticipantDetail | null> {
  try {
    const res = await axios.get(`${API_URL}/participants/${id}`);
    return res.data?.data as ParticipantDetail;
  } catch (error) {
    console.error(error);
    return null;
  }
}

export async function createParticipant(
  name: string,
  email: string,
  genderId: number,
  birthDate: string,
  phoneNumber: string
) {
  try {
    await axios.post(`${API_URL}/participants/add`, {
      name,
      email,
      genderId,
      birthDate,
      phoneNumber,
    });
  } catch (error) {
    console.error(error);
  }
}

export async function updateParticipant(
  id: number,
  name: string,
  email: string,
  genderId: number,
  birthDate: string,
  phoneNumber: string
) {
  try {
    await axios.patch(`${API_URL}/participants/${id}`, {
      name,
      email,
      genderId,
      birthDate,
      phoneNumber,
    });
  } catch (error) {
    console.error(error);
  }
}

export async function deleteParticipant(id: number) {
  try {
    await axios.delete(`${API_URL}/participants/${id}`);
  } catch (error) {
    console.error(error);
  }
}
