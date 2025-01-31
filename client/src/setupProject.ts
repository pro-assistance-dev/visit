import User from '@/classes/User';
import { servicesComponents } from 'sprof';
import { App } from 'vue';
import { Router as R } from 'vue-router';

export default function SetupProject(app: App, router: R) {
  AuthFormSettings[AuthFormStatuses.Login].title = 'Логин';
  AuthFormSettings[AuthFormStatuses.Login].buttons = [new AuthButtonLogin()];
  HttpClient.Setup(import.meta.env.VITE_APP_BASE_URL, import.meta.env.VITE_APP_API_V1, import.meta.env.VITE_APP_API_HOST);
  PF.H.Router.SetRouter(router);
  servicesComponents.install(app);
  Static.Setup(import.meta.env.VITE_APP_STATIC_URL);
  PHelp.Auth.Status.SetUserConstructor(User);
}
