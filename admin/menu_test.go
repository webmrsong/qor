package admin

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/qor/qor/resource"
)

func TestMenu(t *testing.T) {
	var menus []*Menu
	res1 := &Resource{
		Resource: resource.Resource{Name: "res1"},
		Config:   &Config{Menu: []string{"menu1"}},
	}
	res2 := &Resource{
		Resource: resource.Resource{Name: "res2"},
		Config:   &Config{Menu: []string{"menu1"}},
	}
	res3 := &Resource{
		Resource: resource.Resource{Name: "res3"},
		Config:   &Config{Menu: []string{"menu1", "menu1-1"}},
	}
	res4 := &Resource{
		Resource: resource.Resource{Name: "res4"},
		Config:   &Config{Menu: []string{"menu2"}},
	}
	res5 := &Resource{
		Resource: resource.Resource{Name: "res5"},
		Config:   &Config{},
	}
	res6 := &Resource{
		Resource: resource.Resource{Name: "res6"},
		Config:   &Config{Menu: []string{"menu1", "menu1-2"}},
	}
	res7 := &Resource{
		Resource: resource.Resource{Name: "res7"},
		Config:   &Config{Menu: []string{"menu1", "menu1-1", "menu1-1-1"}},
	}

	menus = appendMenu(menus, res7.Config.Menu, res7)
	menus = appendMenu(menus, res1.Config.Menu, res1)
	menus = appendMenu(menus, res2.Config.Menu, res2)
	menus = appendMenu(menus, res3.Config.Menu, res3)
	menus = appendMenu(menus, res4.Config.Menu, res4)
	menus = appendMenu(menus, res5.Config.Menu, res5)
	menus = appendMenu(menus, res6.Config.Menu, res6)
	prefixMenuLinks(menus, "/admin")

	expect := []*Menu{
		&Menu{Name: "menu1", SubMenus: []*Menu{
			&Menu{Name: "menu1-1", SubMenus: []*Menu{
				&Menu{Name: "menu1-1-1", SubMenus: []*Menu{
					&Menu{Name: res7.Name, rawPath: "res7", Link: "/admin/res7"},
				}},
				&Menu{Name: res3.Name, rawPath: "res3", Link: "/admin/res3"},
			}},
			&Menu{Name: res1.Name, rawPath: "res1", Link: "/admin/res1"},
			&Menu{Name: res2.Name, rawPath: "res2", Link: "/admin/res2"},
			&Menu{Name: "menu1-2", SubMenus: []*Menu{
				&Menu{Name: res6.Name, rawPath: "res6", Link: "/admin/res6"},
			}},
		}},
		&Menu{Name: "menu2", SubMenus: []*Menu{
			&Menu{Name: res4.Name, rawPath: "res4", Link: "/admin/res4"},
		}},
		&Menu{Name: res5.Name, rawPath: "res5", Link: "/admin/res5"},
	}

	if !reflect.DeepEqual(expect, menus) {
		g, err := json.MarshalIndent(menus, "", "  ")
		if err != nil {
			t.Error(err)
		}
		w, err := json.MarshalIndent(expect, "", "  ")
		if err != nil {
			t.Error(err)
		}
		t.Errorf("add menu errors: got %s; expect %s", g, w)
	}

	menu := getMenu(menus, "res6")
	if menu == nil {
		t.Error("failed to get menu")
	} else if menu.Name != "res6" {
		t.Error("failed to get correct menu")
	}
}
