package main

type string_list []string

func (arr string_list) indexOf(target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

const infinity = 99999999999

func main() {
	commit1 := string_list{"commit1 message", "foo bar", "hot dog", "bye bye", "ok", "gilbert", "hot dog"}
	commit2 := string_list{"ok", "commit2 message", "hot dog", "foo bar", "bye bye"}

	commit1_matches := []int{}

	for _, line := range commit1 {
		match_index := commit2.indexOf(line)
		commit1_matches = append(commit1_matches, match_index)
	}

	commit2_matches := []int{}

	for _, line := range commit2 {
		match_index := commit1.indexOf(line)
		commit2_matches = append(commit2_matches, match_index)
	}

	// fmt.Println("commit1_matches", commit1_matches)
	// fmt.Println("commit2_matches", commit2_matches)

	commit1_upto := 0
	commit2_upto := 0

	var commit1_path_follow_difference int
	var commit2_path_follow_difference int

	// this is how many lines are considered to be changed when following the commit 1 follow path
	if commit1_matches[commit1_upto] == -1 {
		commit1_path_follow_difference = infinity
	} else {
		commit1_path_follow_difference = commit1_matches[commit1_upto] - commit2_upto

	}

	if commit2_matches[commit2_upto] == -1 {
		commit2_path_follow_difference = infinity
	} else {
		commit2_path_follow_difference = commit2_matches[commit2_upto] - commit1_upto

	}

	if commit1_path_follow_difference < commit2_path_follow_difference {
		commit2
	}

}
