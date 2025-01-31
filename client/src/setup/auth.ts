import { AuthButtonLogin, AuthChunks } from 'sprof';

export default function Setup() {
  // AuthFormSettings[AuthFormChunks.Login].buttons = [new AuthButtonLogin()];
  //
  PF.Auth.Form.Settings[AuthChunks.Login].buttons = [new AuthButtonLogin()];
  // PF.Auth.Status.SetUserConstructor(User);
}
