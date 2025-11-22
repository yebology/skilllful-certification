import axios from "axios";
import { API_URL } from "../env/env";
import type { ClassParticipant, ParticipantClass } from "../types/type";

export async function getClassParticipant(
  classId: number
): Promise<ClassParticipant[]> {
  try {
    const res = await axios.get(API_URL + "/enrollments/classes/" + classId);
    console.log(res.data);
    return (res.data.data as ClassParticipant[]) || [];
  } catch (error) {
    console.error(error);
    return [];
  }
}

export async function getParticipantClass(
  participantId: number
): Promise<ParticipantClass[]> {
  try {
    const res = await axios.get(
      API_URL + "/enrollments/participants/" + participantId
    );
    console.log(res.data);
    return (res.data.data as ParticipantClass[]) || [];
  } catch (error) {
    console.error(error);
    return [];
  }
}

export async function assignParticipant(
  participantId: number,
  classId: number
) {
  try {
    await axios.post(API_URL + "/enrollments/add", {
      classId: classId,
      participantId: participantId,
    });
  } catch (error) {
    console.error(error);
  }
}

export async function deleteClassParticipant(id: number) {
  try {
    await axios.delete(`${API_URL}/enrollments/${id}`);
  } catch (error) {
    console.error(error);
  }
}
