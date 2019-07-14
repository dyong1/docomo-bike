package getstation

import (
	"testing"
)

func TestExtractBikeID(t *testing.T) {
	type testScenario struct {
		line string
		want string
	}
	for scenarioName, scenario := range map[string]testScenario{
		"Includes bike ID": testScenario{
			line: `<a id="cycBtnTab_3" class="cycle_list_btn ui-btn ui-icon-redcarat ui-btn-icon-right ui-nodisc-icon" href="javascript:doubleDisableTab(4); tab_TYO_11686.submit();">TYO7576</a>`,
			want: "TYO_11686",
		},
		"Form input line": testScenario{
			line: `<input type="hidden" name="CycleID" value="11686">`,
			want: "",
		},
	} {
		t.Run(scenarioName, func(t *testing.T) {
			if got, want := extractBikeID(scenario.line), scenario.want; got != want {
				t.Errorf("Expect %s and %s are the same", got, want)
			}
		})
	}
}
