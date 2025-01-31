import { AuthButtonLogin, AuthChunks } from 'sprof';
import User from '@/classes/User';

export default function Setup() {
  // AuthFormSettings[AuthFormChunks.Login].buttons = [new AuthButtonLogin()];
  //
  PF.Auth.Form.Settings[AuthChunks.Login].buttons = [new AuthButtonLogin()];
  PF.Auth.Status.SetUserConstructor(User);
  // PF.Auth.Status.SetUserConstructor(User);
}
