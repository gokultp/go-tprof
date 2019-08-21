package parser

const Res = `
?       bitbucket.org/exotel/iam/gandalf/cmd/http       [no test files]
?       bitbucket.org/exotel/iam/gandalf/cmd/http/handlers      [no test files]
?       bitbucket.org/exotel/iam/gandalf/cmd/http/middlewares   [no test files]
?       bitbucket.org/exotel/iam/gandalf/cmd/tmongo     [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/checks        [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/checks/auth   [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/checks/throttle       [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/configmanager [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/core/entities [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/core/permissions      [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/core/principals       [no test files]
=== RUN   TestCreate
=== RUN   TestCreate/should_throw_error_if_the_given_entity_is_not_defined
=== RUN   TestCreate/should_throw_error_if_the_given_entity's_permission_is_not_matching
=== RUN   TestCreate/should_throw_error_if_the_given_entity's_realm_is_not_matching
=== RUN   TestCreate/should_throw_internal_server_error_if_if_there_is_a_db_error_while_fetching_entities
=== RUN   TestCreate/should_not_throw_any_errors_for_valid_input
--- PASS: TestCreate (0.00s)
    --- PASS: TestCreate/should_throw_error_if_the_given_entity_is_not_defined (0.00s)
    --- PASS: TestCreate/should_throw_error_if_the_given_entity's_permission_is_not_matching (0.00s)
    --- PASS: TestCreate/should_throw_error_if_the_given_entity's_realm_is_not_matching (0.00s)
    --- PASS: TestCreate/should_throw_internal_server_error_if_if_there_is_a_db_error_while_fetching_entities (0.00s)
    --- PASS: TestCreate/should_not_throw_any_errors_for_valid_input (0.00s)
=== RUN   TestDelete
=== RUN   TestDelete/should_throw_not_found_if_enitity_with_sid_is_not_there
=== RUN   TestDelete/should_throw_internal_server_error_if_there_is_a_db_err
=== RUN   TestDelete/should_not_throw_any_errors_if_deletion_happened_properly
--- PASS: TestDelete (0.00s)
    --- PASS: TestDelete/should_throw_not_found_if_enitity_with_sid_is_not_there (0.00s)
    --- PASS: TestDelete/should_throw_internal_server_error_if_there_is_a_db_err (0.00s)
    --- PASS: TestDelete/should_not_throw_any_errors_if_deletion_happened_properly (0.00s)
=== RUN   TestGet
=== RUN   TestGet/should_throw_not_found_if_enitity_with_sid_is_not_there
=== RUN   TestGet/should_throw_internal_server_error_if_there_is_a_db_err
=== RUN   TestGet/should_not_throw_any_errors_if_deletion_happened_properly
--- FAIL: TestGet (0.00s)
    --- FAIL: TestGet/should_throw_not_found_if_enitity_with_sid_is_not_there (0.00s)
        get_test.go:58: Get() got1 = [Not Found]:[Not Found] - [] [], want [Not Found]:[n] - [] []
    --- PASS: TestGet/should_throw_internal_server_error_if_there_is_a_db_err (0.00s)
    --- PASS: TestGet/should_not_throw_any_errors_if_deletion_happened_properly (0.00s)
=== RUN   Test_isSubset
=== RUN   Test_isSubset/should_return_true_for_A=[1,2,3]_and_B=[1,2]
=== RUN   Test_isSubset/should_return_false_for_A=[1,2,3]_and_B=[1,2,_4]
=== RUN   Test_isSubset/should_return_false_for_A=[1,2]_and_B=[1,2,_4]
--- PASS: Test_isSubset (0.00s)
    --- PASS: Test_isSubset/should_return_true_for_A=[1,2,3]_and_B=[1,2] (0.00s)
    --- PASS: Test_isSubset/should_return_false_for_A=[1,2,3]_and_B=[1,2,_4] (0.00s)
    --- PASS: Test_isSubset/should_return_false_for_A=[1,2]_and_B=[1,2,_4] (0.00s)
=== RUN   TestGetEntityNames
=== RUN   TestGetEntityNames/should_return_empty_array_if_input_is_empty
=== RUN   TestGetEntityNames/should_return_empty_array_if_entity_is_null
=== RUN   TestGetEntityNames/should_return_entities__if_entity_is_not_null
--- PASS: TestGetEntityNames (0.00s)
    --- PASS: TestGetEntityNames/should_return_empty_array_if_input_is_empty (0.00s)
    --- PASS: TestGetEntityNames/should_return_empty_array_if_entity_is_null (0.00s)
    --- PASS: TestGetEntityNames/should_return_entities__if_entity_is_not_null (0.00s)
=== RUN   TestUpdate
=== RUN   TestUpdate/should_throw_error_if_the_given_entity_is_not_defined
=== RUN   TestUpdate/should_throw_error_if_the_given_entity's_permission_is_not_matching
=== RUN   TestUpdate/should_throw_error_if_the_given_entity's_realm_is_not_matching
=== RUN   TestUpdate/should_throw_internal_server_error_if_if_there_is_a_db_error_while_fetching_entities
=== RUN   TestUpdate/should_throw_404_if_role_with_given_sid_is_not_found
=== RUN   TestUpdate/should_throw_internal_server_err_if__there_is_a_db_err_while_updating_role
=== RUN   TestUpdate/should_throw_internal_server_err_if_there_is_a_db_err_while_finding_role
=== RUN   TestUpdate/should_not_throw_any_errors_if_there_are_no_db_and_validation_errors
--- PASS: TestUpdate (0.00s)
    --- PASS: TestUpdate/should_throw_error_if_the_given_entity_is_not_defined (0.00s)
    --- PASS: TestUpdate/should_throw_error_if_the_given_entity's_permission_is_not_matching (0.00s)
    --- PASS: TestUpdate/should_throw_error_if_the_given_entity's_realm_is_not_matching (0.00s)
    --- PASS: TestUpdate/should_throw_internal_server_error_if_if_there_is_a_db_error_while_fetching_entities (0.00s)
    --- PASS: TestUpdate/should_throw_404_if_role_with_given_sid_is_not_found (0.00s)
    --- PASS: TestUpdate/should_throw_internal_server_err_if__there_is_a_db_err_while_updating_role (0.00s)
    --- PASS: TestUpdate/should_throw_internal_server_err_if_there_is_a_db_err_while_finding_role (0.00s)
    --- PASS: TestUpdate/should_not_throw_any_errors_if_there_are_no_db_and_validation_errors (0.00s)
FAIL
FAIL    bitbucket.org/exotel/iam/gandalf/internal/core/roles    0.007s
?       bitbucket.org/exotel/iam/gandalf/internal/db    [no test files]
?       bitbucket.org/exotel/iam/gandalf/internal/models        [no test files]
?       bitbucket.org/exotel/iam/gandalf/pkg/client     [no test files]`
