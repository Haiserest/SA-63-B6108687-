import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

//const username = { givenuser: login}
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบประวัติการรักษา' };
 
 return (
   <Page theme={pageTheme.service}>
     <Header
       title={`Welcome ${profile.givenName || 'ระบบประวัติการรักษา'}`}
       subtitle={``}
     >
    {/*       
       {username.givenuser}
    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
           <Button variant="contained" color="secondary" href="#outlined-buttons">
          LogOut
        </Button> 
*/}
     </Header>
     
     <Content>
       <ContentHeader title="ระบบประวัติ">    

         <Link component={RouterLink} to="/create">
           <Button variant="outlined" color="primary">
             Add Data
           </Button>
           </Link> 
&nbsp;&nbsp;&nbsp;&nbsp;
            <Link component={RouterLink} to="/">
           <Button variant="outlined" color="secondary">
             LOG OUT
           </Button>
       </Link> 
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
