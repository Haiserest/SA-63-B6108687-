import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Create from './components/Create';
//import login from './components/Login';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
   // router.registerRoute('/', login);
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/create', Create);
  },
});
 
