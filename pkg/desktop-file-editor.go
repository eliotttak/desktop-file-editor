package dfe

type StrKeyValuePair struct {
	Value    string
	Modifier string
}

type StrSliceKeyValuePair struct {
	Value    []string
	Modifier string
}

type BoolKeyValuePair struct {
	Value    bool
	Modifier string
}

type NumericKeyValuePair struct {
	Value    float64
	Modifier string
}

// Source https://specifications.freedesktop.org/desktop-entry/latest/recognized-keys.html
type DesktopEntry struct {
	Type                  StrKeyValuePair
	Version               StrKeyValuePair
	Name                  StrKeyValuePair
	GenericName           StrKeyValuePair
	NoDisplay             BoolKeyValuePair
	Comment               StrKeyValuePair
	Icon                  StrKeyValuePair
	Hidden                BoolKeyValuePair
	OnlyShowIn, NotShowIn StrSliceKeyValuePair
	DBusActivatable       BoolKeyValuePair
	TryExec               StrKeyValuePair
	Exec                  StrKeyValuePair
	Path                  StrKeyValuePair
	Terminal              BoolKeyValuePair
	Actions               StrSliceKeyValuePair
	MimeType              StrSliceKeyValuePair
	Categories            StrSliceKeyValuePair
	Implements            StrSliceKeyValuePair
	Keywords              StrSliceKeyValuePair
	StartupNotify         BoolKeyValuePair
	StartupWMClass        StrKeyValuePair
	URL                   StrKeyValuePair
	PrefersNonDefaultGPU  BoolKeyValuePair
	SingleMainWindow      BoolKeyValuePair
}

type DesktopAction struct {
	// The ID key is not a real key of a desktop action, but is the action ID defined like [Desktop
	// Action %s] (where %s is the action ID) and should be present in [DesktopEntry.Actions.Value]
	ID   string
	Name StrKeyValuePair
	Icon StrKeyValuePair
	Exec StrKeyValuePair
}

type DesktopFile struct {
	DesktopEntry   DesktopEntry
	DesktopActions []DesktopAction
}

func (df *DesktopFile) AddAction(da DesktopAction) {
	df.DesktopActions = append(df.DesktopActions, da)
	df.DesktopEntry.Actions.Value = append(df.DesktopEntry.Actions.Value, da.ID)
}
