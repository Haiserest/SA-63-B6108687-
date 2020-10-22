import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntTreatment } from '../../api/models/EntTreatment';
import moment from 'moment';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 //const [users, setUsers] = useState(Array);
 const [loading, setLoading] = useState(true);
 const [treatments, setTreatments] = useState<EntTreatment[]>([]);
 
 useEffect(() => {
   const getTreatments = async () => {
     const res = await api.listTreatment({ limit: 10, offset: 0 });
     setLoading(false);
     setTreatments(res);
   };
   getTreatments();
 }, [loading]);
 
 const deleteTreatments = async (id: number) => {
   const res = await api.deleteTreatment({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับที่</TableCell>
           <TableCell align="center">ชื่อ-นามสกุล</TableCell>
           <TableCell align="center">อายุ</TableCell>
           <TableCell align="center">อาการ</TableCell>
           <TableCell align="center">สถานะ</TableCell>
           <TableCell align="center">วันที่การนัดหมายครั้งต่อไป</TableCell>
           <TableCell align="center">แพทย์</TableCell>
           <TableCell align="center">delete</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {treatments.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.patientID?.patientName}</TableCell>
             <TableCell align="center">{item.edges?.patientID?.patientAge}</TableCell>
             <TableCell align="center">{item.symptom}</TableCell>
             <TableCell align="center">{item.edges?.statusID?.status}</TableCell>
             <TableCell align="center">{moment(item.dateAppoint).format("DD-MM-YYYY")}</TableCell>
             <TableCell align="center">{item.edges?.doctorID?.doctorName}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteTreatments(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
             
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
