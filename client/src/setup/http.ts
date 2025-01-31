export default function Setup() {
  PF.H.HttpClient.Setup(import.meta.env.VITE_APP_BASE_URL, import.meta.env.VITE_APP_API_V1, import.meta.env.VITE_APP_API_HOST);
}
