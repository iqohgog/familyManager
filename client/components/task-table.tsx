import { useState, useEffect } from "react";
import {
  useReactTable,
  getCoreRowModel,
  flexRender,
  ColumnDef,
} from "@tanstack/react-table";
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";

// Интерфейс задачи
interface Task {
  ID: string;
  Name: string;
  Description: string;
  AssigneeID: string;
  Priority: string;
  CreatorID: string;
}

// Определение колонок таблицы
const columns: ColumnDef<Task>[] = [
  { accessorKey: "ID", header: "ID" },
  { accessorKey: "Name", header: "Название" },
  { accessorKey: "Description", header: "Описание" },
  { accessorKey: "AssigneeID", header: "ID Исполнителя" },
  { accessorKey: "Priority", header: "Приоритет" },
  { accessorKey: "CreatorID", header: "ID Создателя" },
];

function TaskTable() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isFormOpen, setIsFormOpen] = useState(false); // Состояние для управления формой
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    assignee_id: "",
    priority: "low",
  });

  // Загрузка задач с сервера
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("Токен не найден");
      setLoading(false);
      return;
    }

    fetch("http://localhost:8080/task", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((response) => {
        if (!response.ok) throw new Error("Ошибка загрузки задач");
        return response.json();
      })
      .then((data) => {
        setTasks(data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error.message);
        setLoading(false);
      });
  }, []);

  // Обработчик изменения полей формы
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  // Обработчик изменения приоритета
  const handlePriorityChange = (value: string) => {
    setFormData({ ...formData, priority: value });
  };

  // Обработчик отправки формы
  // Обработчик отправки формы
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem("token");
    if (!token) {
      setError("Токен не найден");
      return;
    }

    const body = {
      name: formData.name,
      description: formData.description,
      assignee_id: formData.assignee_id || undefined,
      priority: formData.priority,
    };

    try {
      const response = await fetch("http://localhost:8080/task", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(body),
      });

      if (!response.ok) throw new Error("Ошибка добавления задачи");
      await response.json(); // Получаем данные новой задачи (можно не использовать)
      setFormData({
        name: "",
        description: "",
        assignee_id: "",
        priority: "low",
      }); // Сбрасываем форму
      setIsFormOpen(false); // Закрываем форму
      window.location.reload(); // Перезагружаем страницу
    } catch (error) {
      setError(error.message);
    }
  };

  const handleCancel = () => {
    setFormData({
      name: "",
      description: "",
      assignee_id: "",
      priority: "low",
    });
    setIsFormOpen(false);
  };

  // Настройка таблицы
  const table = useReactTable({
    data: tasks,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div>
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <TableHead key={header.id}>
                  {flexRender(
                    header.column.columnDef.header,
                    header.getContext()
                  )}
                </TableHead>
              ))}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.map((row) => (
            <TableRow key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>

      {/* Кнопка для открытия формы */}
      <Button onClick={() => setIsFormOpen(true)} className="mt-4">
        Добавить задание
      </Button>

      {/* Форма (показывается только если isFormOpen === true) */}
      {isFormOpen && (
        <form onSubmit={handleSubmit} className="mt-4 space-y-4">
          <Input
            name="name"
            placeholder="Название задания"
            value={formData.name}
            onChange={handleInputChange}
            required
          />
          <Input
            name="description"
            placeholder="Описание"
            value={formData.description}
            onChange={handleInputChange}
            required
          />
          <Input
            name="assignee_id"
            placeholder="ID Исполнителя"
            value={formData.assignee_id}
            onChange={handleInputChange}
          />
          <Select onValueChange={handlePriorityChange} defaultValue="low">
            <SelectTrigger>
              <SelectValue placeholder="Приоритет" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="low">Низкий</SelectItem>
              <SelectItem value="medium">Средний</SelectItem>
              <SelectItem value="high">Высокий</SelectItem>
            </SelectContent>
          </Select>
          <div className="flex space-x-4">
            <Button type="submit">Сохранить</Button>
            <Button type="button" variant="outline" onClick={handleCancel}>
              Отмена
            </Button>
          </div>
        </form>
      )}
    </div>
  );
}

export default TaskTable;
