import * as React from 'react';
import { ToDo } from './ToDoList';
import Checkbox from '@material-ui/core/Checkbox';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import Switch from '@material-ui/core/Switch';

interface ToDoListItemProps {
    todo: ToDo;
    active: boolean;
    handleToggleComplete(todoId: string): () => void;
    handleToggleActive(todoId: string): () => void;
};

class ToDoListItem extends React.Component<ToDoListItemProps, {}> {
    render() {
        const { todo, active, handleToggleComplete, handleToggleActive } = this.props;

        return (
            <ListItem>
                <Checkbox
                    checked={todo.complete}
                    tabIndex={-1}
                    onClick={handleToggleComplete(todo.id)}
                />
                <ListItemText primary={todo.text} />
                <ListItemSecondaryAction>
                    <Switch
                        onChange={handleToggleActive(todo.id)}
                        checked={active && !todo.complete}
                        disabled={todo.complete}
                    />
                </ListItemSecondaryAction>
            </ListItem>
        );
    }
}

export default ToDoListItem;
