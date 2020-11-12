import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
//import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);

export default function Create() {
    
 const classes = useStyles();
 const profile = { givenName: 'ระบบประวัติการรักษา' };
 const api = new DefaultApi();
 
 //const [user, setUser] = useState(initialUserState);
 const [users, setUsers] = useState<EntUser[]>([]);
 //const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [loading, setLoading] = useState(true);

 //const [treatment, setTreatment] = useState(String);
 const [userid, setUserid] = useState(Number);

 useEffect(() => {
   const getUsers = async () => {
     const u = await api.listUser({ limit: 10, offset: 0 });
     setLoading(false);
     setUsers(u);
   };
   getUsers();

 }, [loading]);

 const doctor_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setUserid(event.target.value as number);
 };
 
React.useEffect(() => {
  localStorage.setItem("doctordata",JSON.stringify(userid));
});

const login = async () =>{ 
  if (userid != null){
  setAlert(true);
} else {
  setAlert(false);
}
};

 return (
   <Page theme={pageTheme.service}>
     <Header
       title={`Welcome ${profile.givenName || 'ระบบประว้ติ'}`}
     //  subtitle = {`${username.givenuser}`}
     ></Header>
     <Content>
       <ContentHeader title="ระบบลงทะเบียนประวัติการรักษา">
{/* 
       <Button variant="outlined" color="secondary" href="#outlined-buttons">
          LogOut
        </Button> */}

         {/* {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                ✔️ LOG IN COMPLETE : 😁 😆 🤩!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ❌ CAN'T LOG IN : 🥺 😵 😱!
               </Alert>
             )}
           </div>
         ) : null} */}
         
       </ContentHeader>

       <div className={classes.root}>
         <form noValidate autoComplete="off" >
         <div>
<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="doctor_id-label">DoctorName</InputLabel>
            <Select
              labelId="doctor_id-label"
              label="Doctor"
              id="doctor_id"
              value={userid}
              onChange={doctor_id_handleChange}
              style = {{width: 600}}
              required = {true}
            >
              {users.map((item:EntUser)=>
              <MenuItem value={item.id}>{item.doctorName}</MenuItem>)}
            </Select>
 </FormControl>
 </div>
<center>
           <div className={classes.margin}>
              
             <Button
              onClick={() => {
                login();
              }}
               //style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/welcome"
               variant="contained"
             >
               LOG IN
             </Button>
             
           </div>
</center>
         </form>
       </div>
     </Content>
   </Page>
 );
}
