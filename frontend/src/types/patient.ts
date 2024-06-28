export type Patient = {
  uuid: string;
  phone: string;
  status: string;
  keywords: string[];
};

export type Message = {
  id: number;
  patient_id: string;
  content: string;
  created_at: string;
  loading?: boolean;
  audio?: string;
};
