import Auth from './auth';
import Static from './static';
import Http from './http';
import Router from './router';
import { Router as R } from 'vue-router';
import { servicesComponents } from 'sprof';
import { App } from 'vue';

export default function Setup(app: App, router: R) {
  Auth();
  Static();
  Http();
  Router(router);
  servicesComponents.install(app);
}
