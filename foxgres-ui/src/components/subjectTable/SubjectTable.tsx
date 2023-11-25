import { FC } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from "@mui/material";

interface TableProps {
  disciplines: Array<string>;
  marks: Record<string, number>;
}

const getColorForMark = (mark: number) => {
  switch (mark) {
    case 5:
      return "#29cf00";
    case 4:
      return "#C9FFA0";
    case 3:
      return "#feffa3";
    case 2:
      return "#ffdca6";
    default:
      return "#ff0800";
  }
};

const SubjectTable: FC<TableProps> = ({ disciplines, marks }) => {
  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Дисциплина</TableCell>
            <TableCell align="right">Оценка</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {disciplines.map((discipline) => (
            <TableRow key={discipline}>
              <TableCell component="th" scope="row">
                {discipline}
              </TableCell>
              <TableCell
                align="right"
                sx={{
                  backgroundColor: getColorForMark(marks[discipline]),
                  fontWeight: "bold",
                }}
              >
                {marks[discipline]}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default SubjectTable;
