import axios from "axios";
import { API_URL } from "../env/env";
import type { Class, ClassDetail } from "../types/type";

export async function getAllClass(): Promise<Class[]> {
  try {
    const res = await axios.get(API_URL + "/classes");
    return res.data.data as Class[];
  } catch (error) {
    console.error(error);
    return []; // kembalikan array kosong kalau error
  }
}

export async function getClassDetail(id: number): Promise<ClassDetail | null> {
  try {
    const res = await axios.get(`${API_URL}/classes/${id}`);
    return res.data?.data as ClassDetail;
  } catch (error) {
    console.error(error);
    return null; // <- return null supaya type aman
  }
}

export async function createClasss(
  name: string,
  categoryId: number,
  description: string,
  instructor: string
) {
  try {
    await axios.post(`${API_URL}/classes/add`, {
      name,
      categoryId,
      description,
      instructor,
    });
  } catch (error) {
    console.error(error);
  }
}

export async function updateClasss(
  id: number,
  name: string,
  categoryId: number,
  description: string,
  instructor: string
) {
  try {
    await axios.patch(`${API_URL}/classes/${id}`, {
      name,
      categoryId,
      description,
      instructor,
    });
  } catch (error) {
    console.error(error);
  }
}

export async function deleteClass(id: number) {
  try {
    await axios.delete(`${API_URL}/classes/${id}`);
  } catch (error) {
    console.error(error);
  }
}
