import { API_BASE_URL } from "src/config";

export const fetchGrades = async (studentId: string) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/marks?student_id=${studentId}`,
    );
    if (!response.ok) {
      throw new Error("Проблема с получением данных: " + response.statusText);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Ошибка при выполнении запроса:", error);
    throw error;
  }
};