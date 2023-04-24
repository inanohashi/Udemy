import React, { useState } from 'react';
import './style.css';
import { InputTodo } from './components/inputTodo';
import { IncompleteTodos } from './components/IncompleteTodos'
import { CompleteTodos } from './components/CompleteTodos'

export const App = () => {
    const [todoText, setTodoText] = useState([]);
    const [incompleteTodos, setIncompeteTodos] = useState(["aaaaa", "iiiii"]);
    const [completeTodos, setCompeteTodos] = useState(["uuuuu"]);

    const onChangeTodoText = (event) => setTodoText(event.target.value);

    const onClickAdd = () => {
        if (todoText === "") return;
        const newTodos = [...incompleteTodos, todoText];
        setIncompeteTodos(newTodos);
        setTodoText("");
    };

    const onClickDelete = (index) => {
        const newTodos = [...incompleteTodos];
        newTodos.splice(index, 1);
        setIncompeteTodos(newTodos);
    };

    const onClickComplete = (index) => {
        const newIncompleteTodos = [...incompleteTodos];
        newIncompleteTodos.splice(index, 1);
             
        const newCompleteTodos = [...completeTodos, incompleteTodos[index]];
        setIncompeteTodos(newIncompleteTodos);
        setCompeteTodos(newCompleteTodos);
    };

    const onClickBack = (index) => {
        const newCompleteTodos = [...completeTodos];
        newCompleteTodos.splice(index, 1);

        const newIncompleteTodos = [...incompleteTodos, completeTodos[index]];
        setCompeteTodos(newCompleteTodos);
        setIncompeteTodos(newIncompleteTodos);
    }

    return (
        <>
            <InputTodo todoText={todoText} onChange={onChangeTodoText} onClick={onClickAdd} disabled={incompleteTodos.length >= 5} />
            {incompleteTodos.length >= 5 && <p style={{ color: "red" }}>登録できるtodoは5個までだよ〜</p>}
            <IncompleteTodos todos={incompleteTodos} onClickComplete={onClickComplete} onClickDelete={onClickDelete} />
            <CompleteTodos todos={completeTodos} onClickBack={onClickBack} />
        </>
    );
};