package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/artistomin/planner"
)

type todoRepo struct {
	db         string
	collection string
	session    *mgo.Session
}

func (r *todoRepo) FindByID(id planner.UUID) (*planner.Todo, error) {
	session := r.session.Copy()
	defer session.Close()

	collection := session.DB(r.db).C(r.collection)

	var todo *planner.Todo
	err := collection.FindId(id).One(&todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepo) FindAll() ([]*planner.Todo, error) {
	session := r.session.Copy()
	defer session.Close()
	collection := session.DB(r.db).C(r.collection)

	var todos []*planner.Todo
	err := collection.Find(nil).All(&todos)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepo) Create(todo *planner.Todo) (planner.UUID, error) {
	session := r.session.Copy()
	defer session.Close()
	collection := session.DB(r.db).C(r.collection)

	id := bson.NewObjectId()
	_, err := collection.UpsertId(id, todo)

	if err != nil {
		return "", err
	}

	return planner.UUID(id.Hex()), nil
}

func (r *todoRepo) Update(id planner.UUID, todo *planner.Todo) error {
	session := r.session.Copy()
	defer session.Close()
	collection := session.DB(r.db).C(r.collection)

	err := collection.UpdateId(id, todo)

	if err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) DeleteByID(id planner.UUID) error {
	session := r.session.Copy()
	defer session.Close()
	collection := session.DB(r.db).C(r.collection)

	err := collection.RemoveId(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *todoRepo) Exists(id planner.UUID) (bool, error) {
	session := r.session.Copy()
	defer session.Close()
	collection := session.DB(r.db).C(r.collection)

	objID := bson.ObjectIdHex(string(id))
	err := collection.FindId(objID).Select(bson.M{"_id": 1}).Limit(1).One(nil)

	if err != nil && err == mgo.ErrNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// NewTodoRepository returns a new instance of a Mongo todo repository
func NewTodoRepository(db, collection string, session *mgo.Session) planner.TodoRepo {
	return &todoRepo{
		db:         db,
		collection: collection,
		session:    session,
	}
}
